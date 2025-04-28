// Package main
package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/Asac2142/go-nave-deriva/cmd/api/routes"
)

func main() {
	addr := flag.String("addr", ":3030", "HTTP network address")
	flag.Parse()

	logHandler := slog.NewTextHandler(os.Stdout, nil)
	logger := slog.New(logHandler)

	server := &http.Server{
		Addr:     *addr,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		Handler:  routes.Routes(logger),
	}

	logger.Info("Starting server", "addr", *addr)
	err := server.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
