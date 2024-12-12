import type { CreateSessionDTO, Session } from "@/dto/session.dto"

export interface SessionRepository {
  getSessionByID(sessionID: string): Promise<Session | null>
  insertSession(session: CreateSessionDTO): Promise<void>
  deleteSession(sessionID: string): Promise<void>
  setSessionExpiry(sessionID: string, newExpiry: Date): Promise<void>
}
