package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"net/url"
	"github.com/rs/cors"
)

var authKey string

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	authKey = r.URL.Query().Get("code")	// retrieve the authentication key for the use found in the parameters of the URL
	http.Redirect(w, r, "http://localhost:3000/home", http.StatusFound)
	fmt.Printf("Authentication successful\nAuth Key: %s", authKey)
}

type IncomingData struct { 
	CodeVerifier string `json:"codeVerifier"`
}

func requestAccessToken(someData string) {
	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("code", authKey)
	body.Set("redirect_uri", "http://localhost:8080/callback")
	body.Set("client_id", "98fc1b94f1e445cebcfe067a505598ba")
	body.Set("code_verifier", someData)
	spotifyUrl := "https://accounts.spotify.com/api/token"

	
	req, err := http.PostForm(spotifyUrl, body)
	if err != nil {
		fmt.Println("Error sending a request: ", err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		fmt.Println("HTTP status code:", req.StatusCode)
		return
	}

	var data map[string]interface{}
	responseBody, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	if err := json.Unmarshal(responseBody, &data); err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return
	}
	access_token, ok := data["access_token"].(string)
	if !ok {
		fmt.Println("Invalid access token format in response")
		return
	}
	fmt.Println("Access Token:", access_token)
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
		requestAccessToken(data.CodeVerifier)
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