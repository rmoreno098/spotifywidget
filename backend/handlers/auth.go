package handlers

import (
	"encoding/json"
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

	jwt, err := h.JWT.GenerateToken(user)
	if err != nil {
		slog.Error("Error generating token", "error", err)
		http.Error(w, "Error generating token", 2)
	}

	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    jwt,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)

	// Redirect user to dashboard
	http.Redirect(w, r, "http://localhost:3000/dashboard", http.StatusFound)
}

func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	cookie, err := r.Cookie("auth_token")
	if err != nil {
		slog.Error("Cookie not found", "error", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenStr := cookie.Value
	claims, err := h.JWT.ParseToken(tokenStr)
	if err != nil {
		slog.Error("Error parsing token", "error", err, "token", tokenStr)
		http.Error(w, "Error parsing token", http.StatusUnauthorized)
		return
	}

	user := map[string]string{
		"id":    claims.Sub,
		"name":  claims.Name,
		"email": claims.Email,
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		slog.Error("Error encoding user", "error", err)
		http.Error(w, "Error encoding user", http.StatusInternalServerError)
	}
}
