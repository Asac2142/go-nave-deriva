// Package routes handles routes for the API.
package routes

import (
	"log/slog"
	"net/http"

	"github.com/Asac2142/go-nave-deriva/cmd/api/handlers"
	"github.com/julienschmidt/httprouter"
)

// Routes handles API defined routes.
func Routes(logger *slog.Logger) http.Handler {
	router := httprouter.New()
	handlerNave := handlers.NewNaveLogger(logger)

	router.HandlerFunc(http.MethodGet, "/status", handlerNave.StatusHandler)
	router.HandlerFunc(http.MethodGet, "/repair-bay", handlerNave.RepairBayHandler)
	router.HandlerFunc(http.MethodPost, "/teapot", handlerNave.TeaPotHandler)

	return router
}
