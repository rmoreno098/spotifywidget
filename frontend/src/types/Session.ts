// src/types/Session.ts
export interface User {
  id: string;
  name: string;
  email: string;
}

export interface SessionState {
  isAuthenticated: boolean;
  isLoading: boolean;
  user?: User;
  error?: string;
}

export interface SessionContextType {
  sessionState: SessionState;
  validateSession: () => Promise<void>;
  logout: () => void;
}
