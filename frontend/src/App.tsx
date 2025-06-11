import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import HomePage from './pages/Home';
import DashboardPage from './pages/dashboard';
import PlaylistPage from './pages/playlist';

function App() {
  return (
    <Router>
        <Routes>
          <Route
            path="/" 
            element={ <HomePage />} 
          />
          <Route 
            path="/dashboard" 
            element={ <DashboardPage />} 
          />
          <Route
            path="/playlist/:userId/:playlistId/:playlistName"
            Component={PlaylistPage}
          />
        </Routes>
    </Router>
  );
}

export default App;
