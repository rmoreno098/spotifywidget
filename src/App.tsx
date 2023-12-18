import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import WelcomePage from './components/welcome';
import PlaylistPage from './components/playlist';
import DashboardPage from './components/dashboard';

function App() {
  return (
    <Router>
        <Routes>
          <Route
            path="/" 
            element={ <WelcomePage />} 
          />
          <Route 
            path="/dashboard" 
            element={ <DashboardPage />} 
          />
          <Route
            path="/playlist/:playlistId"
            Component={PlaylistPage}
          />
        </Routes>
    </Router>
  );
}

export default App;