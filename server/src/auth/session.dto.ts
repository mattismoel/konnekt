import type { UserDTO } from "./user.dto";

export type Session = {
  id: string;
  userID: number;
  expiresAt: Date;
}

export type CreateSessionDTO = {
  id: string;
  userID: number;
  expiresAt: Date;
}

export type SessionValidationResult =
  | { session: Session; user: UserDTO }
  | { session: null; user: null }
