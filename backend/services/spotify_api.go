package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"spotify-widget/backend/models"
	"strings"
)

var CLIENT_ID *string
var CLIENT_SECRET *string

func accessToken(code string) (*string, *string, error) {
	params := url.Values{
		"client_id":    {*CLIENT_ID},
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {"http://localhost:8080/api/v1/callback"},
	}
	payload := strings.NewReader(params.Encode())

	resp, err := http.Post("https://accounts.spotify.com/api/token",
		"application/x-www-form-urlencoded",
		payload)
	if err != nil {
		slog.Error("Error attempting to get Authentication token: %s", err.Error())
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body: %s", err.Error())
		return nil, nil, err
	}
	// Print or store the access token (response handling)
	var x models.Token
	err = json.Unmarshal(responseBody, &x)
	if err != nil {
		slog.Error("Error unmarshaling access token response: %s", err.Error())
		return nil, nil, err
	}

	return &x.AccessToken, &x.RefreshToken, nil
}

func profile(token string) (*models.UserProfile, error) {
	url := "https://api.spotify.com/v1/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Error fetching profile from Spotify API: %s", err.Error())
		return nil, err
	}

	// parse the response body and only return the user's id and display name
	var user models.UserProfile
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return &user, nil
}

func tracks(token string, id string) (*http.Response, error) {
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
