package main

import (
	"fmt"
	// "math/rand"
	// "time"
	// "crypto/sha256"
	// "encoding/base64"
	// "bytes"
	"encoding/json"
	"net/http"
	"github.com/rs/cors"
	// "net/url"
	// "github.com/gorilla/mux"
)

// This handler will recieve the url from Spotify's API, indicating that the user has been successfully authenticated.
// The url will contain parameters which are parsed, and is how the Autherization Code is recieved.
// The server then sends a message to the frontend, notifying the status of the authentication.
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	authKey := r.URL.Query().Get("code")	// retrieve the authentication key for the use found in the parameters of the URL
	http.Redirect(w, r, "http://localhost:3000/home", http.StatusFound)
	fmt.Printf("Authentication successful\nAuth Key: %s", authKey)
}

type IncomingData struct { 
	CodeVerifier string `json:"codeVerifier"`
}

func incomingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var data IncomingData
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Failed to parse JSON data :(", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received codeVerifier: %s\n", data.CodeVerifier)
	} else {
		http.Error(w, "Invalid request method :/", http.StatusMethodNotAllowed)
	}
}


func main() {
	corsHandler := cors.Default()

	println("Server is now runnning on port 8080!")
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(callbackHandler)).ServeHTTP(w, r)
    })

	http.HandleFunc("/incoming", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(incomingHandler)).ServeHTTP(w, r)
    })
	
	http.ListenAndServe(":8080", nil)
}