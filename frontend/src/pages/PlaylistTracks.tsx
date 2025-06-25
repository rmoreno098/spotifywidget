import { useState, useEffect } from "react";
import { useSessionGuard } from "../hooks/useSessionGuard";
import { getPlaylistTracks } from "../api/Playlists";
import type { PlaylistTracksResponse } from "../api/Playlists";
import { Link, useParams } from "react-router-dom";
import type { PlaylistTracks } from "../types/Spotify";
import { ArrowLeft, Clock, Headphones } from "lucide-react";
import { ErrorComponent, LoadingComponent } from "../components/Common";

const formatDuration = (ms: number): string => {
  const minutes = Math.floor(ms / 60000);
  const seconds = Math.floor((ms % 60000) / 1000);
  return `${minutes}:${seconds.toString().padStart(2, "0")}`;
};

export default function PlaylistTracks() {
  const [data, setData] = useState<PlaylistTracksResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const { playlistId } = useParams<{ playlistId: string }>();
  const { isAuthenticated } = useSessionGuard();

  useEffect(() => {
    const fetchData = async () => {
      try {
        setLoading(true);
        const playlistTracks = await getPlaylistTracks(playlistId || "");
        setData(playlistTracks);
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

  if (loading) {
    <LoadingComponent />;
  }

  if (error) {
    ErrorComponent(error);
    return null;
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
              <h1 className="text-2xl md:text-3xl font-bold">Your Tracks</h1>
              <p className="text-gray-400 text-sm md:text-base">
                {data?.total || 0} songs
              </p>
            </div>
          </div>
        </div>
      </div>

      {/* Content */}
      <div className="max-w-7xl mx-auto px-4 md:px-8 py-6 md:py-8">
        {data?.items && data.items.length > 0 ? (
          <div className="space-y-2">
            {/* Table Header - Hidden on mobile */}
            <div className="hidden md:grid grid-cols-12 gap-4 px-4 py-2 text-gray-400 text-sm font-medium border-b border-gray-800">
              <div className="col-span-1 text-center">#</div>
              <div className="col-span-5">TITLE</div>
              <div className="col-span-3">ALBUM</div>
              <div className="col-span-2">POPULARITY</div>
              <div className="col-span-1 text-center">
                <Clock size={16} className="mx-auto" />
              </div>
            </div>

            {/* Track List */}
            {data.items.map(({ track }, index) => (
              <div
                key={track.id}
                // onClick={() => handleTrackClick(track.id)}
                className="group grid grid-cols-1 md:grid-cols-12 gap-4 p-4 rounded-lg hover:bg-gray-800 hover:bg-opacity-50 transition-all duration-200 cursor-pointer items-center"
              >
                {/* Mobile Layout */}
                <div className="md:hidden flex items-center space-x-4">
                  {/* Album Art & Play Button */}
                  <div className="relative flex-shrink-0">
                    <div className="w-12 h-12 rounded-md overflow-hidden bg-gradient-to-br from-green-500 to-teal-500">
                      {track.album.images && track.album.images.length > 0 ? (
                        <img
                          src={
                            track.album.images[track.album.images.length - 1]
                              .url
                          }
                          alt={track.album.name}
                          className="w-full h-full object-cover"
                        />
                      ) : (
                        <div className="w-full h-full flex items-center justify-center">
                          <Headphones
                            size={20}
                            className="text-white opacity-70"
                          />
                        </div>
                      )}
                    </div>
                  </div>

                  {/* Track Info */}
                  <div className="flex-1 min-w-0">
                    <h3 className="font-medium text-white truncate group-hover:text-green-400 transition-colors duration-200">
                      {track.name}
                      {track.explicit && (
                        <span className="ml-2 text-xs bg-gray-600 px-1 rounded">
                          E
                        </span>
                      )}
                    </h3>
                    <p className="text-sm text-gray-400 truncate">
                      {track.artists.map((artist) => artist.name).join(", ")} •{" "}
                      {track.album.name}
                    </p>
                    <div className="flex items-center space-x-2 text-xs text-gray-500 mt-1">
                      {/* <span>{formatDuration(track.duration_ms)}</span> */}
                      <span>•</span>
                      <span>{track.popularity}% popular</span>
                    </div>
                  </div>
                </div>

                {/* Desktop Layout */}
                <div className="hidden md:contents">
                  {/* Track Number & Play Button */}
                  <div className="col-span-1 text-center">
                    <div className="group-hover:hidden text-gray-400 text-sm">
                      {index + 1}
                    </div>
                  </div>

                  {/* Title & Artist */}
                  <div className="col-span-5 flex items-center space-x-3 min-w-0">
                    <div className="w-10 h-10 rounded-md overflow-hidden bg-gradient-to-br from-green-500 to-teal-500 flex-shrink-0">
                      {track.album.images && track.album.images.length > 0 ? (
                        <img
                          src={
                            track.album.images[track.album.images.length - 1]
                              .url
                          }
                          alt={track.album.name}
                          className="w-full h-full object-cover"
                        />
                      ) : (
                        <div className="w-full h-full flex items-center justify-center">
                          <Headphones
                            size={16}
                            className="text-white opacity-70"
                          />
                        </div>
                      )}
                    </div>
                    <div className="min-w-0">
                      <h3 className="font-medium text-white truncate group-hover:text-green-400 transition-colors duration-200">
                        {track.name}
                        {track.explicit && (
                          <span className="ml-2 text-xs bg-gray-600 px-1 rounded">
                            E
                          </span>
                        )}
                      </h3>
                      <p className="text-sm text-gray-400 truncate">
                        {track.artists.map((artist, idx) => (
                          <span key={artist.id}>
                            <Link
                              to={`https://open.spotify.com/artist/${artist.id}`}
                              className="hover:text-white hover:underline transition-colors duration-200"
                            >
                              {artist.name}
                            </Link>
                            {idx < track.artists.length - 1 && ", "}
                          </span>
                        ))}
                      </p>
                    </div>
                  </div>

                  {/* Album */}
                  <div className="col-span-3 text-sm text-gray-400 truncate hover:text-white transition-colors duration-200">
                    {track.album.name}
                  </div>

                  {/* Popularity */}
                  <div className="col-span-2">
                    <div className="flex items-center space-x-2">
                      <div className="flex-1 bg-gray-700 rounded-full h-1">
                        <div
                          className="bg-green-500 h-1 rounded-full transition-all duration-300"
                          style={{ width: `${track.popularity}%` }}
                        />
                      </div>
                      <span className="text-xs text-gray-400 w-8">
                        {track.popularity}%
                      </span>
                    </div>
                  </div>

                  {/* Duration */}
                  <div className="col-span-1 flex items-center justify-center space-x-2">
                    <span className="text-sm text-gray-400 min-w-0">
                      {formatDuration(track.duration_ms)}
                    </span>
                  </div>
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="text-center py-16">
            {/*Default View */}
            <Headphones className="mx-auto text-gray-500 mb-4" size={64} />
            <h2 className="text-2xl font-semibold text-gray-300 mb-2">
              No tracks found
            </h2>
            <p className="text-gray-400">
              Start listening to music to see your tracks here!
            </p>
          </div>
        )}
      </div>
    </div>
  );
}
