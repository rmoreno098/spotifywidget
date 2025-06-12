package main

import (
	"log/slog"
	"net/http"
	"spotify-widget-v2/config"
	"spotify-widget-v2/handlers"
	"spotify-widget-v2/routes"
	"spotify-widget-v2/services"
)

func main() {
	slog.Info("Starting...")
	cfg := config.GetConfiguration()

	spotify := services.NewSpotifyService(cfg.SpotifyClientId, cfg.SpotifyClientSecret, cfg.SpotifyRedirectUri)
	jwt := services.NewJwtService(cfg.JWTSecret)

	h := &handlers.Handler{
		Spotify: spotify, JWT: jwt,
	}

	router := routes.GetRouter(h)
	slog.Info("Server listening", "port", cfg.Port)
	panic(http.ListenAndServe(":"+cfg.Port, router))
}
