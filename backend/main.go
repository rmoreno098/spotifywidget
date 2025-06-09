package main

import (
	"log/slog"
	"net/http"
	"spotify-widget/backend/routes"
)

func main() {
	slog.Info("Starting...")
	router := routes.GetRouter()

	panic(http.ListenAndServe(":8080", router))
}
