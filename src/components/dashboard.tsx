import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserProfile, Playlist } from "./types";
import { redirectToAuthCodeFlow, getAccessToken } from "./auth";

function DashboardPage() {
  const clientId = "98fc1b94f1e445cebcfe067a505598ba";
  const navigate = useNavigate();
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

  const playlistClick = async (event: React.MouseEvent<HTMLButtonElement>, playlistId: string) => {
    event.preventDefault();
    console.log("playlist clicked", playlistId);
    navigate(`/playlist/${playlistId}`);
  }

  const spotifyConnect = async (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    if (!code) {
      redirectToAuthCodeFlow(clientId);
    } 
    else {
      const accessToken = await getAccessToken(clientId, code);
      const usr = await fetchProfile(accessToken);
      const usr_playlists = await fetchPlaylists(accessToken);

      if(usr.error || usr_playlists.error) {
        console.error(usr.error.message); // Invalid access token or something similar
        redirectToAuthCodeFlow(clientId);
      } else {
        setProfile(usr);
        setPlaylists(usr_playlists);
        setConnected(true);
      }
    }
  };

  console.log(playlists)
  // {playlists?.items[0].images[0].url}

  return (
    <form className="flex flex-col relative shrink-0 box-border justify-start items-start w-full bg-gray-900 p-9 max-md:mb-9 h-screen">
      <header className="flex flex-col justify-start items-start w-full max-md:gap-[px]">
        <div className="flex flex-col max-w-full self-stretch w-auto max-md:w-auto max-md:self-stretch">
          <header className="gap-5 flex max-md:flex-col max-md:items-stretch max-md:gap-0">
            <div className="flex flex-col items-stretch w-6/12 max-md:w-full max-md:ml-0">
              {connected ? (
                <div className="flex flex-col max-w-full justify-center self-stretch w-full items-start h-full mx-auto max-md:gap-9 max-md:h-auto max-md:grow-0 max-md:mb-9">
                  <div className="flex flex-col justify-center items-start w-auto self-stretch">
                    <h1 className="max-w-[400px] text-white text-4xl tracking-normal text-left mt-2">
                    Hello {profile?.display_name} 👋
                    </h1>                    
                  </div>
                </div>
              ) : (
                <div className="flex flex-col max-w-full justify-center self-stretch w-full items-start h-full mx-auto max-md:gap-9 max-md:h-auto max-md:grow-0 max-md:mb-9">
                  <div className="flex flex-col justify-center items-start w-auto self-stretch">
                    <h1 className="max-w-[350px] text-white text-4xl tracking-normal text-left mt-2">
                      Connect to Spotify
                    </h1>

                    <button
                      className="relative shrink-0 box-border appearance-none text-green-500 bg-green-200 rounded text-center cursor-pointer w-auto self-center mr-auto mt-5 px-6 py-4"
                      onClick={(event)=>spotifyConnect(event)} disabled={connected}
                    >
                      Connect
                    </button>
                  </div>
                </div>
              )
             }
            </div>

            <div className="flex flex-col items-stretch w-6/12 ml-5 max-md:w-full max-md:ml-0">
              <header className="flex flex-row relative shrink-0 box-border w-full justify-between">
                <div className="flex flex-col relative shrink-0 box-border w-full">
                  <header className="gap-5 flex max-md:flex-col max-md:items-stretch max-md:gap-0">
                    <div className="flex flex-col items-stretch w-full max-md:w-full max-md:ml-0">
                        {connected ? ( <></> ) : (
                          <span className="text-white text-4xl tracking-normal text-left mt-2">
                            Connect to see your playlists 🎶
                          </span>
                        )}
                    </div>
                  </header>
                </div>
              </header>
            </div>
          </header>
        </div>
      </header>

      <section className="flex flex-col max-w-full self-stretch w-auto mt-6">
        <header className="gap-5 flex max-md:flex-col max-md:items-stretch max-md:gap-0">
            {playlists?.items.map((playlist) => (
              <div className="flex flex-col items-stretch w-3/12 max-md:w-full max-md:ml-0">
                <button onClick={(event)=>playlistClick(event, playlist.id)}>
                  <img
                    loading="lazy"
                    key={playlist.id}
                    src={playlist.images[0]?.url || "default-playlist-image.jpg"}
                    className="aspect-square object-cover object-bottom w-full shrink-0 box-border min-h-[20px] min-w-[20px] overflow-hidden h-full m-auto rounded-lg max-md:my-6"
                    alt={playlist.name}
                  />
                </button>
                <span className="text-white text-4xl tracking-normal text-center mt-2">
                  {playlist.name}
                </span>
              </div>
            ))}
        </header>
      </section>
    </form>
  );
}

export default DashboardPage;