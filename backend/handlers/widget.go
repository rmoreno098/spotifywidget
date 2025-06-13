package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"spotify-widget-v2/models"
	"spotify-widget-v2/services"
)

func (h *Handler) PlaylistTracks(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(models.AuthContext{Claims: "claims"}).(*services.JwtClaims)

	session, err := h.Redis.GetSession(claims.ID)
	if err != nil {
		slog.Error("Error getting session", "error", err)
		http.Error(w, "Error getting session", http.StatusUnauthorized)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("Error reading body", "error", err)
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}

	var p models.PlaylistItem
	if err := json.Unmarshal(body, &p); err != nil {
		slog.Error("Error unmarshalling body", "error", err)
		http.Error(w, "Error unmarshalling body", http.StatusBadRequest)
		return
	}

	tracks, err := h.Spotify.GetPlaylistTracks(session.AccessToken, p.ID)
	if err != nil {
		slog.Error("Error getting tracks", "error", err)
		http.Error(w, "Error fetching tracks from Spotify", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}

func (h *Handler) Playlists(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(models.AuthContext{Claims: "claims"}).(*services.JwtClaims)

	session, err := h.Redis.GetSession(claims.Sub)
	if err != nil {
		slog.Error("Error getting session", "error", err)
		http.Error(w, "Error getting session", http.StatusUnauthorized)
	}

	playlists, err := h.Spotify.GetPlaylists(session.AccessToken)
	if err != nil {
		slog.Error("Error getting playlists", "error", err)
		http.Error(w, "Error fetching playlists from Spotify", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlists)
}
