export interface PlaylistTracks {
  href: string;
  limit: number;
  next: string;
  offset: number;
  previous: string;
  total: number;
  items: PlaylistTrack[];
}

export interface PlaylistTrack {
  added_at: string;
  // added_by: {
  //   id: string;
  //   name: string;
  // };
  is_local: boolean;
  track: Track;
}

export interface Track {
  album: Album;
  artists: SimplifiedArtist[];
  duration_ms: number;
  explicit: boolean;
  href: string;
  id: string;
  name: string;
  popularity: number;
  preview_url: string;
  track_number: number;
  type: string;
  uri: string;
}

export interface Image {
  url: string;
  height: number;
  width: number;
}

export interface PlaylistItem {
  collaborative: boolean;
  description: string;
  href: string;
  id: string;
  images: Image[];
  name: string;
  public: boolean;
  snapshot_id: string;
  // tracks: {
  //   href: string;
  //   total: number;
  // };
  type: string;
  uri: string;
}

export interface Album {
  album_type: string;
  total_tracks: number;
  href: string;
  id: string;
  images: Image[];
  name: string;
  type: string;
  uri: string;
  artists: SimplifiedArtist[];
}

export interface SimplifiedArtist {
  href: string;
  id: string;
  name: string;
  type: string;
  uri: string;
}
