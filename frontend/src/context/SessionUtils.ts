// src/context/SessionUtils.ts
import { useContext, createContext } from "react";
import type { SessionContextType } from "../types/Session";

export const SessionContext = createContext<SessionContextType | undefined>(
  undefined
);
export const useSession = (): SessionContextType => {
  const context = useContext(SessionContext);
  if (context === undefined) {
    throw new Error("useSession must be used within a SessionProvider");
  }
  return context;
};
