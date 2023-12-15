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

const clientId = "98fc1b94f1e445cebcfe067a505598ba";

function DashboardPage() {
  const params = new URLSearchParams(window.location.search);
  const code = params.get("code");

  const [profile, setProfile] = useState<UserProfile>();
  const [connected, setConnected] = useState(false);

  async function fetchProfile(code: string): Promise<any> {
    const result = await fetch("https://api.spotify.com/v1/me", {
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
        if(usr.error) {
          console.error(usr.error.message); // Invalid access token
          return;
        } else {
          setProfile(usr);
          setConnected(true);
      }
    }
  };

  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      <button onClick={spotifyLogin} disabled={connected}>
        Connect to Spotify
      </button>
      {connected ? <h2>Hello {profile?.display_name}</h2> : <></>}
      <Link to="/">Back to Login</Link>
    </div>
  );
}

export default DashboardPage;