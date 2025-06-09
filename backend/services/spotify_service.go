package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"spotify-widget/server/database"
	"spotify-widget/server/types"
)

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
		slog.Error("Error retreiving user token from DB", err)
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
