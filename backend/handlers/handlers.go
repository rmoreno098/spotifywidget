package handlers

import (
	"spotify-widget-v2/services"
)

type Handler struct {
	Spotify *services.SpotifyService
	JWT     *services.JwtService
	Redis   *services.RedisService
}
