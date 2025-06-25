// src/context/AuthContext.tsx
import type { ReactNode } from "react";
import React, { useEffect, useState } from "react";
import type { User, SessionState } from "../types/Session";
import { SessionContext } from "./SessionUtils";

interface SessionProviderProps {
  children: ReactNode;
}

export const SessionProvider: React.FC<SessionProviderProps> = ({
  children,
}) => {
  const [sessionState, setSessionState] = useState<SessionState>({
    isAuthenticated: false,
    isLoading: true,
  });

  const validateSession = async (): Promise<void> => {
    setSessionState((prev) => ({ ...prev, isLoading: true, error: undefined }));

    try {
      const response = await fetch("http://localhost:8080/api/v1/me", {
        credentials: "include",
      });

      if (response.ok) {
        const user: User = await response.json();
        setSessionState({
          isAuthenticated: true,
          isLoading: false,
          user,
        });
      } else {
        setSessionState({
          isAuthenticated: false,
          isLoading: false,
          error: "Session invalid",
        });
      }
    } catch (error) {
      console.error("Error validating session:", error);
      setSessionState({
        isAuthenticated: false,
        isLoading: false,
        error:
          error instanceof Error ? error.message : "Session validation failed",
      });
    }
  };

  const logout = (): void => {
    // TODO: Implement logout logic
  };

  // Validate session on mount
  useEffect(() => {
    validateSession();
  }, []);

  return (
    <SessionContext.Provider value={{ sessionState, validateSession, logout }}>
      {children}
    </SessionContext.Provider>
  );
};
