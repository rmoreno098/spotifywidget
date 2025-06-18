// src/pages/Dashboard.tsx
import { dashboardItems } from "../assets/DashboardItems";
import { LoadingComponent } from "../components/Common";
import { useSessionGuard } from "../hooks/useSessionGuard";

export default function DashboardPage() {
  const { isLoading, isAuthenticated, user, error } = useSessionGuard();

  if (isLoading) {
    <LoadingComponent />;
  }

  if (!isAuthenticated) {
    window.location.href = "/";
    return null;
  }

  const handleItemClick = (itemId: string) => {
    window.location.href = itemId;
    return null;
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-black text-white p-4 md:p-8">
      {/* Loading State */}
      {error && <div className="error">{error}</div>}

      {/* Header Section */}
      <div className="max-w-7xl mx-auto mb-8 md:mb-12">
        <div className="text-center md:text-left">
          <h1 className="text-3xl md:text-5xl font-bold mb-2">
            Hello, {user?.name}
          </h1>
          <p className="text-gray-300 text-lg md:text-xl">
            What would you like to explore today?
          </p>
        </div>
      </div>

      {/* Grid Section */}
      <div className="max-w-7xl mx-auto">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-6 lg:gap-8">
          {dashboardItems.map((item) => {
            const IconComponent = item.icon;
            return (
              <div
                key={item.id}
                onClick={() => handleItemClick(item.id)}
                className="group relative overflow-hidden rounded-2xl cursor-pointer transform transition-all duration-300 hover:scale-105 hover:-translate-y-2"
              >
                {/* Background Gradient */}
                <div
                  className={`absolute inset-0 bg-gradient-to-br ${item.gradient} group-hover:${item.hoverGradient} transition-all duration-300`}
                />

                {/* Content */}
                <div className="relative p-8 md:p-10 lg:p-12 h-48 md:h-56 lg:h-64 flex flex-col justify-between">
                  {/* Icon */}
                  <div className="flex justify-between items-start">
                    <IconComponent
                      size={48}
                      className="text-white drop-shadow-lg group-hover:scale-110 transition-transform duration-300"
                    />
                    <div className="w-3 h-3 rounded-full bg-white bg-opacity-30 group-hover:bg-opacity-50 transition-all duration-300" />
                  </div>

                  {/* Text Content */}
                  <div className="space-y-2">
                    <h2 className="text-2xl md:text-3xl font-bold text-white drop-shadow-lg">
                      {item.title}
                    </h2>
                    <p className="text-white text-opacity-90 text-sm md:text-base font-medium">
                      {item.description}
                    </p>
                  </div>
                </div>

                {/* Hover Effect Ring */}
                <div className="absolute inset-0 rounded-2xl ring-2 ring-white ring-opacity-0 group-hover:ring-opacity-20 transition-all duration-300" />
              </div>
            );
          })}
        </div>
      </div>

      {/* Footer Spacer */}
      <div className="h-8 md:h-16" />
    </div>
  );
}
