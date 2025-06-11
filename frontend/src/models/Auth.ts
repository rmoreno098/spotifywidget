// src/models/Auth.ts
export type AuthUser = {
  sub: string;
  name: string;
  email: string;
  exp: number;
  iat: number;
};

export type AuthState = {
  isAuthenticated: boolean;
  user: AuthUser | null;
};
