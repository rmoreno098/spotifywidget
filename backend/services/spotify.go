package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"spotify-widget-v2/models"
	"strings"
)

type SpotifyService struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func NewSpotifyService(clientID string, clientSecret string, redirectURI string) *SpotifyService {
	return &SpotifyService{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	}
}

func (s *SpotifyService) GenerateTokens(code string) (*models.Token, error) {
	params := url.Values{
		"code":         {code},
		"redirect_uri": {s.RedirectURI},
		"grant_type":   {"authorization_code"},
	}
	payload := strings.NewReader(params.Encode())

	r, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", payload)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	auth := s.ClientID + ":" + s.ClientSecret
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	r.Header.Add("Authorization", "Basic "+encoded)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var t models.Token
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *SpotifyService) GetUserProfile(token string) (*models.UserProfile, error) {
	endpoint := "https://api.spotify.com/v1/me"
	r, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var u models.UserProfile
	if err := json.NewDecoder(res.Body).Decode(&u); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *SpotifyService) GetPlaylistTracks(token string, playlistID string) (*models.PlaylistTracks, error) {
	endpoint := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks", playlistID)
	r, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var t models.PlaylistTracks
	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return nil, err
	}

	return &t, nil
}

func (s *SpotifyService) GetPlaylists(token string) (*models.Playlists, error) {
	endpoint := "https://api.spotify.com/v1/me/playlists"
	r, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var p models.Playlists
	if err := json.NewDecoder(res.Body).Decode(&p); err != nil {
		return nil, err
	}

	return &p, nil
}
