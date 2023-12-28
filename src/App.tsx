import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import WelcomePage from './components/welcome';
import PlaylistPage from './components/playlist';
import DashboardPage from './components/dashboard';
import AnalyzerPage from './components/analyzer';

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
            path="/playlist/:userId/:playlistId/:playlistName"
            Component={PlaylistPage}
          />
					<Route
						path="/analyzer/:userId"
						Component={AnalyzerPage}
						/>
        </Routes>
    </Router>
  );
}

export default App;
