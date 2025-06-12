package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		slog.Error("Code not provided in callback")
		http.Error(w, "No Authorization Code", 1)
		return
	}

	// Exchange authentication code for an access token
	token, err := h.Spotify.GenerateTokens(code)
	if err != nil {
		slog.Error("Fetching access token error", "error", err)
		http.Error(w, "Error fetching access token", 2)
		return
	}

	// Fetch user profile
	user, err := h.Spotify.GetUserProfile(token.AccessToken)
	if err != nil {
		slog.Error("Error fetching user profile", "error", err)
		http.Error(w, "Error fetching profile", 2)
		return
	}
	slog.Info("Successfully fetched user profile", "user", user)

	// Redirect user to dashboard
	redirectURL := fmt.Sprintf("http://localhost:3000/dashboard")
	http.Redirect(w, r, redirectURL, http.StatusFound)
}
