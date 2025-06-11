// src/App.tsx
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import HomePage from "./pages/Home";
import DashboardPage from "./pages/Dashboard";
import PlaylistPage from "./pages/Playlist";

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/dashboard" element={<DashboardPage />} />
        <Route
          path="/playlist/:userId/:playlistId/:playlistName"
          Component={PlaylistPage}
        />
      </Routes>
    </Router>
  );
}
