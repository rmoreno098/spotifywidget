// components/ProtectedRoute.tsx
import React from "react";
import type { ReactNode } from "react";
import { useSession } from "../context/SessionUtils";
import { LoadingComponent } from "./Common";
import HomePage from "../pages/Home";

interface ProtectedRouteProps {
  children: ReactNode;
  fallback?: ReactNode;
}

export const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children }) => {
  const { sessionState } = useSession();

  if (sessionState.isLoading) {
    return <LoadingComponent />;
  }

  if (!sessionState.isAuthenticated) {
    return <HomePage />;
  }

  return <>{children}</>;
};
