import { useParams } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { Playlist } from './types';
import { getTracks } from './auth';
import Chart from 'chart.js/auto';

const PlaylistPage = () => {

    const { userId, playlistId, playlistName } = useParams();
    const [tracks, setTracks] = useState<Playlist>();
    const [loading, setLoading] = useState(true);
  
    useEffect(() => {
        fetchTracks();
    }, []);

    async function fetchTracks() {
        if (userId && playlistId) { 
            const tracks = await getTracks(userId, playlistId);
            if(tracks === null) { return }
            const dona = handleData(tracks);
            generateChart(dona);
            setTracks(tracks);
            setLoading(false);
        }
    }

    async function generateChart(donut: Record<string, { artist: string, count: number }>) {
        const canvas = document.getElementById('donutChart') as HTMLCanvasElement;
        if(canvas) {
            const existingChart = Chart.getChart(canvas);
            if(existingChart) {
                existingChart.destroy();
            }
            new Chart(canvas, {
                type: 'doughnut',
                data: {
                    labels: Object.values(donut).map(entry => entry.artist),
                    datasets: [{
                        data: Object.values(donut).map(entry => entry.count),
                        backgroundColor: [
                            '#FF6384',
                            '#36A2EB',
                            '#FFCE56',
                            '#8B008B',
                        ],
                    }],
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                },
            });
        }
    }

    function handleData(data: Playlist) {
        const donut: Record<string, { artist: string, count: number }> = {};
        for (let i = 0; i < data.total; i++) {
            const artistId = data.items[i].track.artists[0].id;    // artist id
            const artist = data.items[i].track.artists[0].name;   // artist
            if(!donut[artistId]) {
                donut[artistId] = { artist: artist, count: 1 };
            } else {
                donut[artistId].count++;
            }
        }
        return donut;
    }

    return (
        <div className='bg-gray-900'>
            <h1 className="text-white text-4xl  tracking-normal text-center mb-5">{playlistName}</h1>
            <header className="h-[800px] w-full flex items-center justify-center">
                <canvas id="donutChart" className="my-4 mx-auto" width="400" height="400"></canvas>
                {loading ? ( 
                    <h1 className="text-white text-4xl  tracking-normal text-center mb-5">Loading...</h1>
                ) : (
                    <canvas id="donutChart" className="my-4 mx-auto" width="400" height="400"></canvas>
                )}
            </header>

            <h1 className="text-white text-4xl  tracking-normal text-center mb-5">Tracks</h1>
            <div className="grid grid-cols-3 gap-2.5">
                {tracks ? 
                    tracks.items.map((obj) => (
                        <div key={obj.track.id} className='flex flex-col w-full mb-2.5'>     
                            <button onClick={()=>{window.open(obj.track.external_urls.spotify, '_blank')}}>
                                <img className="block w-full rounded-lg shadow-lg border-2 border-black hover:border-green-500"
                                    key={obj.track.id} 
                                    src={obj.track.album.images[0].url} 
                                    alt={obj.track.name}
                                />
                            </button>
                            <span className="text-white text-xl tracking-normal text-center">
                                {obj.track.name}
                            </span>
                        </div>
                    ))
                    : 
                    <h1>Loading tracks...</h1>
                }
            </div>
        </div>
    );
};

export default PlaylistPage;