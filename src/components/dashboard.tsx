import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { UserProfile, Playlist } from "./types";
import { getPlaylists } from "./auth";

function DashboardPage() {
  const navigate = useNavigate();
  const urlParams = new URLSearchParams(window.location.search);

  const userId = urlParams.get('userId');
  const userName = urlParams.get('name');
  const [profile, setProfile] = useState<UserProfile["display_name"]>();
  const [playlists, setPlaylists] = useState<Playlist>();
  const [connected, setConnected] = useState(false);

  useEffect(() => {
    if (userName !== null) {
      if(sessionStorage.getItem(userName)) {
        const userPlaylists = sessionStorage.getItem(userName);
        if (userPlaylists !== null ) {
          const parsed = JSON.parse(userPlaylists);
          console.log(parsed);
          setProfile(userName);
          setPlaylists(parsed);
          setConnected(true);
        }
      }
    }
  }, []);

  const playlistClick = async (event: React.MouseEvent<HTMLButtonElement>, playlistId: string, playlistName: string) => {
    event.preventDefault();
    navigate(`/playlist/${userId}/${playlistId}/${playlistName}`);
  }
	const analyzerClick = async (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    navigate(`/analyzer/${userId}`);
  }

  const spotifyConnect = async (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    if (userId === null || userName === null) {
      console.log("Error getting user info")
      return;
    } 
    // get user's playlists from Spotify
    const userPlaylists = await getPlaylists(userId);
    if (userPlaylists === null ) {
      console.log("Error getting user playlists")
      return;
    }
    setProfile(userName);
    setPlaylists(userPlaylists);
    setConnected(true);
    sessionStorage.setItem(userName, JSON.stringify(userPlaylists));
  };
  
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
                    Hello {profile} 👋
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
										<button
                      className="relative shrink-0 box-border appearance-none text-green-500 bg-green-200 rounded text-center cursor-pointer w-auto self-center mr-auto mt-5 px-6 py-4"
                      onClick={(event)=>analyzerClick(event)} disabled={connected}
                    >
                      Analyze Spotify
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
                        {connected ? ( null ) : (
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
              <div key={playlist.id} className="flex flex-col items-stretch w-3/12 max-md:w-full max-md:ml-0">
                <button onClick={(event)=>playlistClick(event, playlist.id, playlist.name)}>
                  <img
                    loading="lazy"
                    key={playlist.id}
                    src={playlist.images[0]?.url}
                    className="aspect-square object-cover object-bottom w-full shrink-0 box-border border-black hover:border-green-500 overflow-hidden h-full m-auto rounded-lg max-md:my-6"
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
