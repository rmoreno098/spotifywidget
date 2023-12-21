package types

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

type PlaylistResp struct {
	UserId string `json:"user_id"`
}

type TrackResp struct { 
	UserId string `json:"user_id"`
	PlaylistId string `json:"playlist_id"`
}