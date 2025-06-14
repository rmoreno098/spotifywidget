// src/components/Common.tsx
import { Music } from "lucide-react";

export function LoadingComponent() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-black text-white">
      <div className="flex items-center justify-center h-screen">
        <div className="flex flex-col items-center space-y-4">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-green-500"></div>
          <p className="text-gray-300">Loading...</p>
        </div>
      </div>
    </div>
  );
}

export function ErrorComponent(error: string) {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-black text-white">
      <div className="flex items-center justify-center h-screen">
        <div className="text-center space-y-3">
          <Music className="mx-auto text-gray-500" size={48} />
          <h2 className="text-xl font-semibold text-gray-300">
            Something went wrong
          </h2>
          <p className="text-gray-400">{error}</p>
        </div>
      </div>
    </div>
  );
}
