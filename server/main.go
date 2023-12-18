package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/rs/cors"
	"log"
	"spotify-widget/server/types"
)
var verifier string

func fetchProfile(token string) (*http.Response, error) {
	url := "https://api.spotify.com/v1/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func verifierHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)	// read the body of the request
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	var x defs.VerResp	// create a variable of type verResp
    err = json.Unmarshal(body, &x)	// store body into x
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")	// retrieve the code found in the parameters of the callback URL
	if code == "" {
		log.Println("Authentication failed")
		return
	}

	token := getAccessToken(code, verifier)
	if token == "error" {
		log.Println("Authentication failed")
		return
	}

	resp, err := fetchProfile(token)
	if err != nil {
		log.Println("Error fetching profile")
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	log.Println(resp)

	rawJSON, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Error reading Spotify response", http.StatusInternalServerError)
        return
    }

    // Send the raw JSON data to the frontend
    w.Header().Set("Content-Type", "application/json")
    w.Write(rawJSON)
}

func getAccessToken(code string, verifier string) string {
	params := url.Values{
		"client_id":     {"98fc1b94f1e445cebcfe067a505598ba"},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {"http://localhost:8080/callback"},
		"code_verifier": {verifier},
	}
	payload := strings.NewReader(params.Encode())

	resp, err := http.Post("https://accounts.spotify.com/api/token", 
						   "application/x-www-form-urlencoded",
						   payload)
	if err != nil {
		log.Println(err)
		return "error"
	} else {
		defer resp.Body.Close()

		// Read the response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return "error"
		}

		// Print or store the access token (response handling)
		var x defs.TokenResp
		err = json.Unmarshal(responseBody, &x)
		if err != nil {
			log.Println(err)
			return "error"
		}

		return x.AccessToken
	}
}

func main() {
	corsHandler := cors.Default()

	log.Println("Server is now runnning on port 8080!")

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(callbackHandler)).ServeHTTP(w, r)
    })

	http.HandleFunc("/verifier", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(verifierHandler)).ServeHTTP(w, r)
    })

	http.ListenAndServe(":8080", nil)
}
