const { VITE_SPOTIFY_REDIRECT_URI } = import.meta.env;

export async function redirectToAuthCodeFlow(clientId: string) {
  const params = new URLSearchParams();
  params.append("client_id", clientId);
  params.append("response_type", "code");
  params.append("redirect_uri", VITE_SPOTIFY_REDIRECT_URI);
  params.append(
    "scope",
    "user-read-private user-read-email user-top-read playlist-read-private",
  );

  window.location.href = `https://accounts.spotify.com/authorize?${params.toString()}`;
}
