// Package handlers handles incoming HTTP requests.
package handlers

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"
	"strconv"

	errs "github.com/Asac2142/go-nave-deriva/internal/errors"
	"github.com/Asac2142/go-nave-deriva/internal/models"
	"github.com/Asac2142/go-nave-deriva/internal/store"
)

// NaveHandler handler struct definition.
type NaveHandler struct {
	logger *slog.Logger
}

// NewNaveLogger returns a new instance of NaveHandler.
func NewNaveLogger(logger *slog.Logger) *NaveHandler {
	return &NaveHandler{logger: logger}
}

// StatusHandler handles status request calls.
func (nh *NaveHandler) StatusHandler(w http.ResponseWriter, _ *http.Request) {
	respStatus, err := store.Read[models.StatusInfo]()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := respStatus.Status
	response := models.NewResponse(&key)
	jsn, err := json.Marshal(response)
	if err != nil {
		nh.logger.Error(err.Error())
		http.Error(w, errs.INTERNALSERVICE, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsn)
}

// RepairBayHandler handles repair bay request calls.
func (nh *NaveHandler) RepairBayHandler(w http.ResponseWriter, _ *http.Request) {
	statusInfo, err := store.Read[models.StatusInfo]()
	if err != nil {
		nh.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value := models.DamageSchema[statusInfo.Status]

	template, err := template.ParseFiles("./public/html/index.page.tmpl")
	if err != nil {
		nh.logger.Error(err.Error())
		http.Error(w, errs.INTERNALSERVICE, http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, value)
	if err != nil {
		nh.logger.Error(err.Error())
		http.Error(w, errs.INTERNALSERVICE, http.StatusInternalServerError)
		return
	}
}

// TeaPotHandler handles tea pot request calls.
func (nh *NaveHandler) TeaPotHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}

// PhaseChangeDiagram handles volume liquid & vapor reading calls.
func (nh *NaveHandler) PhaseChangeDiagram(w http.ResponseWriter, r *http.Request) {
	const v1 = 0.0015
	const v2 = 30.00
	const p1 = 0.05
	const p2 = 10
	const vc = 0.0035

	param := r.URL.Query().Get("pressure")
	pressure, err := strconv.ParseFloat(param, 64)
	if err != nil {
		nh.logger.Error(err.Error())
		http.Error(w, errs.INVALIDPRESSURE, http.StatusBadRequest)
		return
	}

	if pressure == p1 || pressure == p2 {
		var resp *models.VolumeResponse

		if pressure == p1 {
			resp = models.NewVolumeResponse(v1, v2)
		} else {
			resp = models.NewVolumeResponse(vc, vc)
		}

		jsn, err := json.Marshal(resp)
		if err != nil {
			nh.logger.Error(err.Error())
			http.Error(w, errs.INTERNALSERVICE, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsn)
		return
	}

	if pressure > p1 && pressure < p2 {
		fraction := (pressure - p1) / (p2 - p1)
		volumeLiquid := v1 + fraction*(vc-v1)
		volumeVapor := v2 - fraction*(v2-vc)
		resp := models.NewVolumeResponse(volumeLiquid, volumeVapor)
		jsn, err := json.Marshal(resp)
		if err != nil {
			nh.logger.Error(err.Error())
			http.Error(w, errs.INTERNALSERVICE, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsn)
		return
	}
}
