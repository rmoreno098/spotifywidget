package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"spotify-widget/server/types"
	"strings"
	"github.com/rs/cors"
	"spotify-widget/server/database"
)

var verifier string

func fetchProfile(token string) (string, string, error) {
	url := "https://api.spotify.com/v1/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	// parse the response body and only return the user's id and display name
	var x types.UserProfile
	err = json.NewDecoder(resp.Body).Decode(&x)
	if err != nil {
		return "", "", err
	}
	resp.Body.Close()

	return x.ID, x.DisplayName, nil
}

func verifierHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)	// read the body of the request
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	var x types.VerResp	// create a variable of type verResp
    err = json.Unmarshal(body, &x)	// store body into x
	if err != nil {
		fmt.Println(err)
		return
	}
	verifier = x.Verifier	// store the verifier into the global variable

	w.WriteHeader(http.StatusOK)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve the code found in the parameters of the callback URL
	code := r.URL.Query().Get("code")	
	if code == "" {
		log.Println("Authentication failed")
		return
	}

	// exchange the code for an access token
	token := getAccessToken(code, verifier)
	if token == "error" {
		log.Println("Authentication failed")
		return
	}

	// fetch the user's id and display name
	id, name, err := fetchProfile(token)
	if err != nil {
		log.Println(err)
	}

	// store the user's id, name, and token into the database
	err = database.StoreUsrToken(id, name, token)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "http://localhost:5173/", http.StatusNotFound)
		return
	}

	// redirect the user to the dashboard if the token was successfully stored
	redirectURL := fmt.Sprintf("http://localhost:5173/dashboard?userId=%s&name=%s", url.QueryEscape(id), url.QueryEscape(name))
	http.Redirect(w, r, redirectURL, http.StatusFound)
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
		var x types.TokenResp
		err = json.Unmarshal(responseBody, &x)
		if err != nil {
			fmt.Println(err)
			return "error"
		}

		return x.AccessToken
	}
}

func playlistsHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve the user's id from the request
	body, err := io.ReadAll(r.Body)	// read the body of the request
	if err != nil {
		fmt.Println(err)
		return
	}
	var x types.PlaylistResp	// create a variable of type PlaylistResp
    err = json.Unmarshal(body, &x)	// store body into x
	if err != nil {
		fmt.Println(err)
		return
	}
	id := x.UserId

	// retrieve the user's token from the database
	token, err := database.GetUsrToken(id)
	if err != nil {
		log.Println("Error retreiving user token from DB", err)
		http.Error(w, "Error retreiving user token from DB", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// fetch the user's playlists
	resp, err := fetchPlaylists(token)
	if err != nil {
		log.Println("Error fetching playlists from Spotify", err)
		http.Error(w, "Error fetching playlists from Spotify", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// read the response body
	rawJSON, err := io.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, "Error reading Spotify response", http.StatusInternalServerError)
        return
    }

	// write the response body to the client
	w.Header().Set("Content-Type", "application/json")
    w.Write(rawJSON)
}

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve the playlist id from the request
	body, err := io.ReadAll(r.Body)	// read the body of the request
	if err != nil {
		fmt.Println(err)
		return
	}
	var x types.TrackResp	// create a variable of type PlaylistResp
    err = json.Unmarshal(body, &x)	// store body into x
	if err != nil {
		fmt.Println(err)
		return
	}
	user := x.UserId
	playlist := x.PlaylistId

	// retrieve the user's token from the database
	token, err := database.GetUsrToken(user)
	if err != nil {
		log.Println("Error retreiving user token from DB", err)
		http.Error(w, "Error retreiving user token from DB", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// fetch the user's playlists
	resp, err := fetchTracks(token, playlist)
	if err != nil {
		http.Error(w, "Error fetching tracks from Spotify", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// read the response body
	rawJSON, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading Spotify response", http.StatusInternalServerError)
		return
	}

	// write the response body to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(rawJSON)
}

func fetchTracks(token string, id string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", id)
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

func fetchPlaylists(token string) (*http.Response, error) {
	url := "https://api.spotify.com/v1/me/playlists"
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

func main() {
	err := database.InitDB()
	if err != nil {
		log.Println(err)
		return
	}
	defer database.CloseDB()

	corsHandler := cors.Default()

	log.Println("Server is now runnning on port 8080!")

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(callbackHandler)).ServeHTTP(w, r)
    })

	http.HandleFunc("/verifier", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(verifierHandler)).ServeHTTP(w, r)
    })

	http.HandleFunc("/getPlaylists", func(w http.ResponseWriter, r *http.Request) {
		corsHandler.Handler(http.HandlerFunc(playlistsHandler)).ServeHTTP(w, r)
	})

	http.HandleFunc("/getTracks", func(w http.ResponseWriter, r *http.Request) {
		corsHandler.Handler(http.HandlerFunc(tracksHandler)).ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}