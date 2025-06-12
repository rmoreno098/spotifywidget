package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"spotify-widget-v2/handlers"
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
		slog.Info(fmt.Sprintf("%s %s\n %q", r.Method, r.URL, r.Body))
		next.ServeHTTP(w, r) // Call the next handler
	})
}

func setupRouter(h *handlers.Handler) (*mux.Router, error) {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		message := "Hello, World!"
		_, err := w.Write([]byte(message))
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	}).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/callback", h.Callback).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/v1/tracks", h.PlaylistTracks).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/playlists", h.Playlists).Methods(http.MethodGet)

	return r, nil
}
