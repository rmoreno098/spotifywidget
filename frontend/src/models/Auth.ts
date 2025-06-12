// src/models/Auth.ts
export type User = {
  id: string;
  name: string;
  email: string;
};

export type AuthState = {
  isAuthenticated: boolean;
  user: User | null;
};
