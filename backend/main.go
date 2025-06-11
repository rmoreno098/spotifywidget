package main

import (
	"log/slog"
	"net/http"
	"spotify-widget/backend/config"
	"spotify-widget/backend/routes"
)

func main() {
	slog.Info("Starting...")
	config := config.GetConfiguration()
	router := routes.GetRouter()
	slog.Info(config.Port)

	panic(http.ListenAndServe(":"+config.Port, router))
}
