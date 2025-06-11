package routes

import (
	"fmt"
	"net/http"
	"spotify-widget/backend/services"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router, err := setupRouter()
	if err != nil {
		panic(fmt.Sprintf("Error setting up router %s", err.Error()))
	}
	return router
}

func setupRouter() (*mux.Router, error) {
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
	r.HandleFunc("/api/v1/tracks", services.Tracks).Methods(http.MethodGet)

	return r, nil
}
