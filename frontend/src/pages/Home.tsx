// src/pages/Home.tsx
import { useNavigate } from "react-router-dom";

import loginImage from "../assets/Home.jpeg";
import { useAuth } from "../hooks/useAuth";
import { redirectToAuthCodeFlow } from "../api/auth";

const { VITE_SPOTIFY_CLIENT_ID } = import.meta.env;

export default function HomePage() {
  const clientId = VITE_SPOTIFY_CLIENT_ID;
  const navigate = useNavigate();

  // Check if the user is already authenticated
  const {isAuthenticated} = useAuth();
  if (isAuthenticated) {
    navigate("/dashboard");
    return null;
  }

  async function spotifyConnect() {
    try {
      await redirectToAuthCodeFlow(clientId);
    } catch (error) {
      console.error("Failed to redirect to Spotify auth flow:", error);
      alert(
        "An error occurred while trying to connect to Spotify. Please try again."
      );
    }
  }

  return (
    <div className="min-h-screen bg-gray-900 text-white flex">
      {/* Image Section */}
      <div className="hidden md:flex md:w-3/5 lg:w-1/2">
        <img
          loading="lazy"
          src={loginImage}
          alt="Spotify Welcome Page"
          className="w-full h-full object-cover"
        />
      </div>

      {/* Content Section */}
      <div className="flex-1 flex items-center justify-center p-8 md:w-2/5 lg:w-1/2">
        <div className="w-full max-w-md text-center space-y-8">
          <h1 className="text-3xl md:text-4xl font-bold text-white leading-tight">
            Connect to Spotify
          </h1>

          <button
            onClick={spotifyConnect}
            className="w-full max-w-xs mx-auto bg-green-500 hover:bg-green-600 text-white font-semibold py-3 px-6 rounded-lg text-lg transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 focus:ring-offset-gray-900"
          >
            Get Started
          </button>
        </div>
      </div>
    </div>
  );
}
