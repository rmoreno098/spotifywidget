export async function getTracks(userId: string, playlistId: string) {
  const result = await fetch("http://localhost:8080/getTracks", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ user_id: userId, playlist_id: playlistId }),
  });

  const tracks = await result.json();
  return tracks;
}
