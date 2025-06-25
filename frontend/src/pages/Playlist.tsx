// src/pages/Playlist.tsx
import { useEffect, useState } from "react";
import { ArrowLeft, Play, Music, MoreHorizontal } from "lucide-react";
import { getPlaylists } from "../api/Playlists";
import type { PlaylistsResponse } from "../api/Playlists";
import { ErrorComponent, LoadingComponent } from "../components/Common.tsx";
import { useSessionGuard } from "../hooks/useSessionGuard.ts";
import { Link } from "react-router-dom";

export default function PlaylistPage() {
  const [data, setData] = useState<PlaylistsResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const { isAuthenticated } = useSessionGuard();

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const playlists = await getPlaylists();
        setData(playlists);
      } catch (error: unknown) {
        setError((error as Error)?.message || "An unknown error occurred");
      } finally {
        setLoading(false);
      }
    };

    if (!isAuthenticated) {
      window.location.href = "/";
      return;
    }

    fetchData();
  }, []);

  const handlePlaylistClick = (playlistId: string) => {
    window.location.href = `/playlistracks/${playlistId}`;
    return null;
  };

  if (loading) {
    return <LoadingComponent />;
  }

  if (error) {
    return ErrorComponent(error);
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-black text-white">
      {/* Header */}
      <div className="sticky top-0 z-10 bg-gray-900 bg-opacity-90 backdrop-blur-sm border-b border-gray-800">
        <div className="max-w-7xl mx-auto px-4 md:px-8 py-4">
          <div className="flex items-center space-x-4">
            <Link
              to="/dashboard"
              className="p-2 rounded-full bg-gray-800 hover:bg-gray-700 transition-colors duration-200"
            >
              <ArrowLeft size={20} />
            </Link>
            <div>
              <h1 className="text-2xl md:text-3xl font-bold">Your Playlists</h1>
              <p className="text-gray-400 text-sm md:text-base">
                {data?.total || 0} playlists
              </p>
            </div>
          </div>
        </div>
      </div>

      {/* Content */}
      <div className="max-w-7xl mx-auto px-4 md:px-8 py-6 md:py-8">
        {data?.items && data.items.length > 0 ? (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 md:gap-6">
            {data.items.map((playlist) => (
              <div
                key={playlist.id}
                onClick={() => handlePlaylistClick(playlist.id)}
                className="group bg-gray-800 bg-opacity-50 rounded-xl p-4 hover:bg-gray-700 hover:bg-opacity-70 transition-all duration-300 cursor-pointer transform hover:scale-105 hover:-translate-y-1"
              >
                {/* Playlist Image */}
                <div className="relative group mb-4 aspect-square rounded-lg overflow-hidden bg-gradient-to-br from-purple-500 to-pink-500">
                  {playlist.images && playlist.images.length > 0 ? (
                    <img
                      src={playlist.images[0].url}
                      alt={playlist.name}
                      className="w-full h-full object-cover"
                    />
                  ) : (
                    <div className="w-full h-full flex items-center justify-center">
                      <Music size={48} className="text-white opacity-70" />
                    </div>
                  )}

                  {/* Play Button Overlay */}
                  <div className="absolute inset-0 flex items-center justify-center transition-all duration-300 z-10 bg-black/0 group-hover:bg-black/30">
                    <div className="opacity-0 group-hover:opacity-100 transition-opacity duration-300 pointer-events-auto">
                      <div className="bg-green-500 rounded-full p-3 shadow-lg hover:bg-green-400 transition-colors duration-200">
                        <Play
                          size={20}
                          className="text-black ml-1"
                          fill="currentColor"
                        />
                      </div>
                    </div>
                  </div>
                </div>

                {/* Playlist Info */}
                <div className="space-y-2">
                  <h3 className="font-semibold text-white text-lg line-clamp-2 group-hover:text-green-400 transition-colors duration-200">
                    {playlist.name}
                  </h3>

                  {playlist.description && (
                    <p className="text-gray-400 text-sm line-clamp-2">
                      {playlist.description}
                    </p>
                  )}
                </div>

                {/* Action Menu */}
                <div className="absolute top-6 right-6 opacity-0 group-hover:opacity-100 transition-opacity duration-200">
                  <button className="p-1 rounded-full bg-gray-900 bg-opacity-70 hover:bg-opacity-90 transition-all duration-200">
                    <MoreHorizontal size={16} className="text-white" />
                  </button>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="text-center py-16">
            <Music className="mx-auto text-gray-500 mb-4" size={64} />
            <h2 className="text-2xl font-semibold text-gray-300 mb-2">
              No playlists found
            </h2>
            <p className="text-gray-400">
              Start creating playlists to see them here!
            </p>
          </div>
        )}
      </div>
    </div>
  );
}
