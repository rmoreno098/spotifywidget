package routes

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"spotify-widget-v2/handlers"
	"spotify-widget-v2/models"
)

func GetRouter(h *handlers.Handler) http.Handler {
	router, err := setupRouter(h)
	if err != nil {
		panic(fmt.Sprintf("Error setting up router %s", err.Error()))
	}
	return loggingMiddleware(router)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Incoming request", "method", r.Method, "url", r.URL, "body", r.Body)
		next.ServeHTTP(w, r) // Call the next handler
	})
}

func authMiddleware(h *handlers.Handler) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

			ctx := context.WithValue(r.Context(), models.AuthContext{Claims: "claims"}, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func setupRouter(h *handlers.Handler) (*mux.Router, error) {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/callback", h.Callback).Methods(http.MethodGet)

	// Register protected routes
	protected := api.PathPrefix("/").Subrouter()
	protected.Use(authMiddleware(h))
	protected.HandleFunc("/me", h.Me).Methods(http.MethodGet)
	protected.HandleFunc("/tracks", h.PlaylistTracks).Methods(http.MethodGet)
	protected.HandleFunc("/playlists", h.Playlists).Methods(http.MethodGet)

	return r, nil
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	message := "pong"
	_, err := w.Write([]byte(message))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
