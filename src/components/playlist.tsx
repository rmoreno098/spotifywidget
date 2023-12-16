import { useParams } from 'react-router-dom';

const PlaylistPage = () => {
    
    const { playlistId } = useParams();

    return (
        <h1>This is the Playlist Page for id: {playlistId} !</h1>
    );
};

export default PlaylistPage;