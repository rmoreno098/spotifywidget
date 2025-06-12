export async function getPlaylists(userId: string) {
  const result = await fetch("http://localhost:8080/getPlaylists", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ user_id: userId }),
  });

  const playlists = await result.json();
  return playlists;
}
