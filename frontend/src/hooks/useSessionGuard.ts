// hooks/useSessionGuard.ts
import { useEffect } from "react";
import { useSession } from "../context/SessionUtils";

interface UseSessionGuardOptions {
  redirectTo?: string;
  requireAuth?: boolean;
  onUnauthorized?: () => void;
}

export const useSessionGuard = (options: UseSessionGuardOptions = {}) => {
  const { sessionState, validateSession } = useSession();
  const { redirectTo = "/", requireAuth = true, onUnauthorized } = options;

  useEffect(() => {
    if (sessionState.isLoading) return;

    if (requireAuth && !sessionState.isAuthenticated) {
      if (onUnauthorized) {
        onUnauthorized();
      } else {
        window.location.href = redirectTo;
      }
    }
  }, [
    sessionState.isAuthenticated,
    sessionState.isLoading,
    requireAuth,
    redirectTo,
    onUnauthorized,
  ]);

  return {
    isLoading: sessionState.isLoading,
    isAuthenticated: sessionState.isAuthenticated,
    user: sessionState.user,
    error: sessionState.error,
    retry: validateSession,
  };
};
