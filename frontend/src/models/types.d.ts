export interface UserProfile {
  country: string;
  display_name: string;
  email: string;
  explicit_content: {
    filter_enabled: boolean;
    filter_locked: boolean;
  };
  external_urls: { spotify: string };
  followers: { href: string; total: number };
  href: string;
  id: string;
  images: Image[];
  product: string;
  type: string;
  uri: string;
}
  
export interface Image {
  url: string;
  height: number;
  width: number;
}
  
export interface Playlist {
  href: string;
  items: Item[];
  limit: number;
  next: string;
  offset: number;
  previous: string;
  total: number;
}
  
export interface Item {
  collaborative: boolean;
  description: string;
  external_urls: { spotify: string };
  href: string;
  id: string;
  images: Image[];
  name: string;
  owner: Owner; // do later
  primary_color: string;
  public: boolean;
  snapshot_id: string;
  type: string;
  uri: string;
  track: Track;
}

export interface Track {
  album: Album;
  artists: Artist[];
  available_markets: string[];
  disc_number: number;
  duration_ms: number;
  explicit: boolean;
  external_ids: Externalids;
  external_urls: { spotify: string };
  href: string;
  id: string;
  is_local: boolean;
  name: string;
  popularity: number;
  preview_url: string;
  track_number: number;
  type: string;
  uri: string;
}

export interface Album {
  images: Image[];
  name: string;
}

// not used yet
export interface Artist {
  genres: string[];
  href: string;
  id: string;
  name: string;
  popularity: number;
  uri: string;
}