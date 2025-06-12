// src/context/AuthUtils.ts
import { createContext, useContext } from "react";
import type { AuthState } from "../models/Auth";

export const AuthContext = createContext<AuthState>({
  isAuthenticated: false,
  user: null,
});

export const useAuth = () => useContext(AuthContext);
