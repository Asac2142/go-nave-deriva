// Package handlers handles incoming HTTP requests.
package handlers

import (
	"encoding/json"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/Asac2142/go-nave-deriva/internal/models"
	"github.com/Asac2142/go-nave-deriva/internal/store"
)

const errMsg = "The server encountered a problem and could not process your request"

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
		http.Error(w, errMsg, http.StatusInternalServerError)
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
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	err = template.Execute(w, value)
	if err != nil {
		nh.logger.Error(err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
}

// TeaPotHandler handles tea pot request calls.
func (nh *NaveHandler) TeaPotHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusTeapot)
}
