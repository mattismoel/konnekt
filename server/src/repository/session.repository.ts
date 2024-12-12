import type { CreateSessionDTO, Session } from "@/dto/session.dto"

export interface SessionRepository {
  getByID(sessionID: string): Promise<Session | null>
  insert(session: CreateSessionDTO): Promise<void>
  delete(sessionID: string): Promise<void>
  setExpiry(sessionID: string, newExpiry: Date): Promise<void>
}
