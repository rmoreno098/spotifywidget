package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"spotify-widget/server/database"
	"spotify-widget/server/types"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	// retrieve the code found in the parameters of the callback URL
	code := r.URL.Query().Get("code")
	if code == "" {
		slog.Error("No code found")
		http.Error(w, "No Authorization Code", 1)
		return
	}

	// exchange the code for an access token
	access_token, _, err := accessToken(code)
	if err != nil {
		slog.Error("Fetching access token error")
		http.Error(w, "Error fetching access token", 2)
	}

	// fetch the user's id and display name
	user, err := profile(*access_token)
	if err != nil {
		slog.Error(fmt.Sprintf("Error fetching profile: %s", err.Error()))
	}
	// redirect the user to the dashboard if the token was successfully stored
	redirectURL := fmt.Sprintf("http://localhost:5173/dashboard?userId=%s&name=%s", url.QueryEscape(user.ID), url.QueryEscape(user.DisplayName))
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func Tracks(w http.ResponseWriter, r *http.Request) {
	// retrieve the playlist id from the request
	body, err := io.ReadAll(r.Body) // read the body of the request
	if err != nil {
		fmt.Println(err)
		return
	}
	var x types.TrackResp          // create a variable of type PlaylistResp
	err = json.Unmarshal(body, &x) // store body into x
	if err != nil {
		fmt.Println(err)
		return
	}
	user := x.UserId
	playlist := x.PlaylistId

	// retrieve the user's token from the database
	token, _, err := database.GetUserToken(user)
	if err != nil {
		slog.Error(fmt.Sprintf("Error retreiving user token from DB: %s", err.Error()))
		http.Error(w, "Error retreiving user token from DB", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// fetch the user's playlists
	resp, err := tracks(token, playlist)
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
