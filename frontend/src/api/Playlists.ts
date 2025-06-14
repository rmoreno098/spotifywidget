export const getPlaylists = async () => {
  try {
    const result = await fetch("http://localhost:8080/api/v1/playlists", {
      credentials: "include",
    });

    const playlists = await result.json();
    return playlists;
  } catch (error) {
    console.log("An error occurred fetching playlists:", error);
    return error;
  }
};
