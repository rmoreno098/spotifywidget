// src/context/AuthContext.tsx
import React, { useEffect, useState } from "react";
import type { User, AuthState } from "../models/Auth";
import { AuthContext } from "./AuthUtils";

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [authState, setAuthState] = useState<AuthState>({
    isAuthenticated: false,
    user: null,
  });

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/v1/me", {
          credentials: "include",
        });
        if (res.ok) {
          const data: User = await res.json();
          if (data) {
            setAuthState({ isAuthenticated: true, user: data });
          }
        }
      } catch (err) {
        console.error("Failed to fetch user info:", err);
        window.location.href = "/";
        setAuthState({ isAuthenticated: false, user: null });
      }
    };

    fetchUser();
  }, []);

  return (
    <AuthContext.Provider value={authState}>{children}</AuthContext.Provider>
  );
};
