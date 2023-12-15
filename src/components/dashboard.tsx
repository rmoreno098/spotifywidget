import { Link, useLocation } from "react-router-dom";
import { redirectToAuthCodeFlow, getAccessToken } from "./auth";
import "bootstrap/dist/css/bootstrap.min.css";
import { useEffect, useState } from "react";

interface UserProfile {
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

interface Image {
  url: string;
  height: number;
  width: number;
}

interface Playlist {
  href: string;
  items: Item[];
  limit: number;
  next: string;
  offset: number;
  previous: string;
  total: number;
}

interface Item {
  collaborative: boolean;
  description: string;
  external_urls: { spotify: string };
  href: string;
  id: string;
  images: Image[];
  name: string;
  owner: Owner;
  primary_color: string;
  public: boolean;
  snapshot_id: string;
  tracks: Tracks;
  type: string;
  uri: string;
}

interface Tracks {
  href: string;
  total: number;
}

interface Owner {
  display_name: string;
  external_urls: { spotify: string };
  href: string;
  id: string;
  type: string;
  uri: string;
}

const clientId = "98fc1b94f1e445cebcfe067a505598ba";

function DashboardPage() {
  const params = new URLSearchParams(window.location.search);
  const code = params.get("code");

  const [profile, setProfile] = useState<UserProfile>();
  const [playlists, setPlaylists] = useState<Playlist>();
  const [connected, setConnected] = useState(false);

  async function fetchProfile(code: string): Promise<any> {
    const result = await fetch("https://api.spotify.com/v1/me", {
      method: "GET",
      headers: { Authorization: `Bearer ${code}` },
    });
    return await result.json();
  }

  async function fetchPlaylists(code: string): Promise<any> {
    const result = await fetch("https://api.spotify.com/v1/me/playlists", {
      method: "GET",
      headers: { Authorization: `Bearer ${code}` },
      });
      return await result.json();
  }

  const spotifyLogin = async () => {
    console.log("button clicked !");
    if (!code) {
      redirectToAuthCodeFlow(clientId);
    } else {
        const accessToken = await getAccessToken(clientId, code);
        const usr = await fetchProfile(accessToken);
        const usr_playlists = await fetchPlaylists(accessToken);
        if(usr.error || usr_playlists.error) {
          console.error(usr.error.message); // Invalid access token
          return;
        } else {
          setProfile(usr);
          setPlaylists(usr_playlists);
          setConnected(true);
      }
    }
  };

  function showProfile() {
    return (
      <div>

        <div className="flex">
          <h2>Hello {profile?.display_name}</h2>
          <img src={profile?.images[0].url} alt="profile" />
          <h2>Followers: {profile?.followers.total}</h2>
          <h2>Playlists: {playlists?.total}</h2>
        </div>

      </div>
    );
  }

  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      <button onClick={spotifyLogin} disabled={connected}>
        Connect to Spotify
      </button>
      {connected ? showProfile() : <></>}
      <Link to="/">Back to Login</Link>
    </div>
  );
}

export default DashboardPage;