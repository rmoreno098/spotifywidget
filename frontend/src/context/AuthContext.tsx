// src/context/AuthContext.tsx
import React, { useEffect, useState } from "react";
import { AuthContext } from "./AuthUtils";
import type { User, AuthState } from "../models/Auth";

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [authState, setAuthState] = useState<AuthState>({
    isAuthenticated: false,
    user: null,
    loading: true,
  });

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const res = await fetch("http://localhost:8080/api/v1/me", {
          credentials: "include",
        });
        if (res.ok) {
          const data: User = await res.json();
          setAuthState({ isAuthenticated: true, user: data, loading: false });
        } else {
          setAuthState({ isAuthenticated: false, user: null, loading: false });
        }
      } catch (err) {
        console.error("Failed to fetch user info:", err);
        setAuthState({ isAuthenticated: false, user: null, loading: false });
      }
    };

    fetchUser();
  }, []);

  return (
    <AuthContext.Provider value={authState}>{children}</AuthContext.Provider>
  );
};
