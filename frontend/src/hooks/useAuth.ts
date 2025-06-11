// src/hooks/useAuth.ts
import { useState, useEffect } from "react";
import Cookies from "js-cookie";

import type { AuthUser, AuthState } from "../models/Auth";

const decodeJwt = (token: string): AuthUser | null => {
  try {
    const payload = token.split(".")[1];
    const decoded = atob(payload.replace(/-/g, "+").replace(/_/g, "/"));
    return JSON.parse(decoded);
  } catch (error) {
    console.error("Failed to decode JWT:", error);
    return null;
  }
};

export const useAuth = (): AuthState => {
  const [authState, setAuthState] = useState<AuthState>({
    isAuthenticated: false,
    user: null,
  });

  useEffect(() => {
    const token = Cookies.get("auth_token");
    console.log("Token from cookies:", token);
    if (token) {
      const user = decodeJwt(token);
      setAuthState({ isAuthenticated: !!user, user });
    } else {
      setAuthState({ isAuthenticated: false, user: null });
    }
  }, []);

  console.log("Auth state:", authState);

  return authState;
};
