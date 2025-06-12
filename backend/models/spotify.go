package models

type Track struct {
	Album       Album              `json:"album"`
	Artists     []SimplifiedArtist `json:"artists"`
	DurationMS  int                `json:"duration_ms"`
	Explicit    bool               `json:"explicit"`
	Href        string             `json:"href"`
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Popularity  int                `json:"popularity"`
	TrackNumber int                `json:"track_number"`
	Type        string             `json:"type"`
	Uri         string             `json:"uri"`
}

//type Playlist struct {
//	ID       string  `json:"id"`
//	Href     string  `json:"href"`
//	Items    []Item  `json:"items"`
//	Limit    int     `json:"limit"`
//	Next     string  `json:"next"`
//	Offset   int     `json:"offset"`
//	Previous string  `json:"previous"`
//	Total    int     `json:"total"`
//	PlaylistTracks   []Track `json:"tracks"`
//}

type PlaylistTracks struct {
	Href     string          `json:"href"`
	Limit    int             `json:"limit"`
	Next     string          `json:"next"`
	Offset   int             `json:"offset"`
	Previous string          `json:"previous"`
	Total    int             `json:"total"`
	Items    []PlaylistTrack `json:"items"`
}

type PlaylistTrack struct {
	AddedAt string `json:"added_at"`
	//AddedBy string `json:"added_by"`
	IsLocal bool  `json:"is_local"`
	Track   Track `json:"track"`
}

type Playlists struct {
	Href     string         `json:"href"`
	Limit    int            `json:"limit"`
	Next     string         `json:"next"`
	Offset   int            `json:"offset"`
	Previous string         `json:"previous"`
	Total    int            `json:"total"`
	Items    []PlaylistItem `json:"items"`
}

type PlaylistItem struct {
	Collaborative bool    `json:"collaborative"`
	Description   string  `json:"description"`
	Href          string  `json:"href"`
	ID            string  `json:"id"`
	Images        []Image `json:"images"`
	Name          string  `json:"name"`
	Public        bool    `json:"public"`
	SnapshotID    string  `json:"snapshot_id"`
	//PlaylistTracks        []Track `json:"tracks"`	// Review
	Type string `json:"type"`
	URI  string `json:"uri"`
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

type Album struct {
	AlbumType   string             `json:"album_type"`
	TotalTracks int                `json:"total_tracks"`
	Href        string             `json:"href"`
	Id          string             `json:"id"`
	Images      []Image            `json:"images"`
	Name        string             `json:"name"`
	Type        string             `json:"type"`
	Uri         string             `json:"uri"`
	Artists     []SimplifiedArtist `json:"artists"`
}

type SimplifiedArtist struct {
	Href string `json:"href"`
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Uri  string `json:"uri"`
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

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}
