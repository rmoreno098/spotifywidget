package main

import (
	"log/slog"
	"net/http"
	"spotify-widget-v2/config"
	"spotify-widget-v2/routes"
)

func main() {
	slog.Info("Starting...")
	config := config.GetConfiguration()
	router := routes.GetRouter()
	slog.Info(config.Port)

	panic(http.ListenAndServe(":"+config.Port, router))
}
