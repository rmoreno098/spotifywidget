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
	redis := services.NewRedisService(cfg.RedisAddr, cfg.RedisPassword)

	h := &handlers.Handler{
		Spotify: spotify, JWT: jwt, Redis: redis,
	}

	router := routes.GetRouter(h)
	slog.Info("Server listening", "port", cfg.Port)
	panic(http.ListenAndServe(":"+cfg.Port, router))
}
