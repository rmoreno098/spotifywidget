// src/App.tsx
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./pages/Home";
import DashboardPage from "./pages/Dashboard";
import PlaylistPage from "./pages/Playlist";
import { ProtectedRoute } from "./components/ProtectedRoute";
import PlaylistTracks from "./pages/PlaylistTracks";

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route
          path="/dashboard"
          element={
            <ProtectedRoute>
              <DashboardPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/playlists"
          element={
            <ProtectedRoute>
              <PlaylistPage />
            </ProtectedRoute>
          }
        />
        <Route
          path="/playlistracks/:playlistId"
          element={
            <ProtectedRoute>
              <PlaylistTracks />
            </ProtectedRoute>
          }
        />
      </Routes>
    </Router>
  );
}
