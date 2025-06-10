package services

import (
	"fmt"
	"net/http"
)

func fetchTracks(token string, id string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
