package models

type Track struct {
	Album    Album    `json:"album"`
	Artists  []Artist `json:"artists"`
	Duration int      `json:"duration_ms"`
	Explicit bool     `json:"explicit"`
	Href     string   `json:"href"`
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Uri      string   `json:"uri"`
}

type TopTracks struct {
	Href     string  `json:"href"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
	Items    []Track `json:"items"`
}

type Artist struct {
	// ExternalURLs ExternalURLs `json:"external_urls"`
	Genres     []string `json:"genres"`
	Href       string   `json:"href"`
	Id         string   `json:"id"`
	Images     []Image  `json:"images"`
	Name       string   `json:"name"`
	Popularity int      `json:"popularity"`
	Type       string   `json:"type"`
	Uri        string   `json:"uri"`
}

type Album struct {
	AlbumType   string  `json:"album_type"`
	TotalTracks int     `json:"total_tracks"`
	HRef        string  `json:"href"`
	Id          string  `json:"id"`
	Images      []Image `json:"images"`
	Name        string  `json:"images"`
}

type Playlists struct {
	Href     string `json:"href"`
	Items    []Item `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type UserProfile struct {
	Country     string  `json:"country"`
	DisplayName string  `json:"display_name"`
	Email       string  `json:"email"`
	Href        string  `json:"href"`
	ID          string  `json:"id"`
	Images      []Image `json:"images"`
	Product     string  `json:"product"`
	Type        string  `json:"type"`
	URI         string  `json:"uri"`
}

type Item struct {
	Collaborative bool    `json:"collaborative"`
	Description   string  `json:"description"`
	Href          string  `json:"href"`
	ID            string  `json:"id"`
	Images        []Image `json:"images"`
	Name          string  `json:"name"`
	PrimaryColor  string  `json:"primary_color"`
	Public        bool    `json:"public"`
	SnapshotID    string  `json:"snapshot_id"`
	Type          string  `json:"type"`
	URI           string  `json:"uri"`
}

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}
