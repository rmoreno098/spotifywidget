// src/assets/DashboardItems.ts
import { Music, Headphones, Mic, Users } from "lucide-react";

export const dashboardItems = [
  {
    id: "playlists",
    title: "Playlists",
    icon: Music,
    description: "Your curated collections",
    gradient: "from-purple-500 to-pink-500",
    hoverGradient: "from-purple-600 to-pink-600",
  },
  {
    id: "tracks",
    title: "Tracks",
    icon: Headphones,
    description: "Your favorite songs",
    gradient: "from-green-500 to-teal-500",
    hoverGradient: "from-green-600 to-teal-600",
  },
  {
    id: "podcasts",
    title: "Podcasts",
    icon: Mic,
    description: "Shows you follow",
    gradient: "from-blue-500 to-indigo-500",
    hoverGradient: "from-blue-600 to-indigo-600",
  },
  {
    id: "artists",
    title: "Artists",
    icon: Users,
    description: "Musicians you love",
    gradient: "from-orange-500 to-red-500",
    hoverGradient: "from-orange-600 to-red-600",
  },
];
