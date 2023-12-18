package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"github.com/rs/cors"
)

type VerResp struct {
	Verifier string `json:"verifier"`
}

type TokenResp struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope string `json:"scope"`
}

type UserProfile struct {
	Country          string            `json:"country"`
	DisplayName      string            `json:"display_name"`
	Email            string            `json:"email"`
	ExplicitContent  ExplicitContent   `json:"explicit_content"`
	ExternalURLs     ExternalURLs      `json:"external_urls"`
	Followers        Followers         `json:"followers"`
	Href             string            `json:"href"`
	ID               string            `json:"id"`
	Images           []Image           `json:"images"`
	Product          string            `json:"product"`
	Type             string            `json:"type"`
	URI              string            `json:"uri"`
}

type ExplicitContent struct {
	FilterEnabled bool `json:"filter_enabled"`
	FilterLocked  bool `json:"filter_locked"`
}

type ExternalURLs struct {
	Spotify string `json:"spotify"`
}

type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Playlist struct {
	Href     string `json:"href"`
	Items    []Item `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type Item struct {
	Collaborative bool          `json:"collaborative"`
	Description   string        `json:"description"`
	ExternalURLs  ExternalURLs  `json:"external_urls"`
	Href          string        `json:"href"`
	ID            string        `json:"id"`
	Images        []Image       `json:"images"`
	Name          string        `json:"name"`
	Owner         Owner         `json:"owner"`
	PrimaryColor  string        `json:"primary_color"`
	Public        bool          `json:"public"`
	SnapshotID    string        `json:"snapshot_id"`
	Tracks        Tracks        `json:"tracks"`
	Type          string        `json:"type"`
	URI           string        `json:"uri"`
}

type Tracks struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Owner struct {
	DisplayName  string       `json:"display_name"`
	ExternalURLs ExternalURLs `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

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
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	var x VerResp	// create a variable of type verResp
    err = json.Unmarshal(body, &x)	// store body into x
	if err != nil {
		fmt.Println(err)
		return
	}
	verifier = x.Verifier	// store the verifier into the global variable

	w.WriteHeader(http.StatusOK)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")	// retrieve the code found in the parameters of the callback URL
	if code == "" {
		fmt.Println("Authentication failed")
		return
	}

	token := getAccessToken(code, verifier)
	if token == "error" {
		fmt.Println("Authentication failed")
		return
	}

	resp, err := fetchProfile(token)
	if err != nil {
		fmt.Println("Error fetching profile")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

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
		fmt.Println(err)
		return "error"
	} else {
		defer resp.Body.Close()

		// Read the response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return "error"
		}

		// Print or store the access token (response handling)
		var x TokenResp
		err = json.Unmarshal(responseBody, &x)
		if err != nil {
			fmt.Println(err)
			return "error"
		}

		return x.AccessToken
	}
}

func main() {
	corsHandler := cors.Default()

	fmt.Println("Server is now runnning on port 8080!")

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(callbackHandler)).ServeHTTP(w, r)
    })

	http.HandleFunc("/verifier", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(verifierHandler)).ServeHTTP(w, r)
    })

	http.ListenAndServe(":8080", nil)
}