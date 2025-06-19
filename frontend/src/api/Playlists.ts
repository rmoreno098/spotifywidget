import type { PlaylistItem, PlaylistTrack } from "../types/Spotify";

export interface PlaylistsResponse {
  items: PlaylistItem[];
  total: number;
}

export interface PlaylistTracksResponse {
  items: PlaylistTrack[];
  total: number;
}

export const getPlaylists = async (): Promise<PlaylistsResponse> => {
  const result = await fetch("http://localhost:8080/api/v1/playlists", {
    credentials: "include",
  });

  const playlists = await result.json();
  return playlists;
};

export const getPlaylistTracks = async (
  playlistId: string,
): Promise<PlaylistTracksResponse> => {
  const result = await fetch(
    `http://localhost:8080/api/v1/playlistracks/${playlistId}`,
    {
      credentials: "include",
    },
  );
  const playlistTracks = await result.json();
  return playlistTracks;
};
