package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"spotify-widget-v2/models"

	"github.com/gorilla/mux"
)

func (h *Handler) PlaylistTracks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playlistID := vars["id"]
	if playlistID == "" {
		http.Error(w, "PlaylistID not found", http.StatusBadRequest)
		return
	}

	session := r.Context().Value(models.AuthContext{Session: "session"}).(*models.Token)

	tracks, err := h.Spotify.GetPlaylistTracks(session.AccessToken, playlistID)
	if err != nil {
		slog.Error("Error getting tracks", "error", err)
		http.Error(w, "Error fetching tracks from Spotify", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}

func (h *Handler) Playlists(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(models.AuthContext{Session: "session"}).(*models.Token)

	playlists, err := h.Spotify.GetPlaylists(session.AccessToken)
	if err != nil {
		slog.Error("Error getting playlists", "error", err)
		http.Error(w, "Error fetching playlists from Spotify", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlists)
}
