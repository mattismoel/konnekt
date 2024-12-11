import type { CreateSessionDTO, Session } from "@/dto/session.dto"
import type { TX } from "./db"
import { sessionsTable } from "./schema/session"
import { eq } from "drizzle-orm"

export const insertSessionTx = async (tx: TX, session: CreateSessionDTO): Promise<void> => {
  await tx
    .insert(sessionsTable)
    .values(session)
}

export const getSessionByIDTx = async (tx: TX, sessionID: string): Promise<Session | null> => {
  const results = await tx
    .select()
    .from(sessionsTable)
    .where(eq(sessionsTable.id, sessionID))

  if (results.length <= 0) {
    return null
  }

  return results[0]
}

export const deleteSessionTx = async (tx: TX, sessionID: string): Promise<void> => {
  await tx
    .delete(sessionsTable)
    .where(eq(sessionsTable.id, sessionID))
}

export const setSessionExpiryTx = async (tx: TX, sessionID: string, newExpiry: Date): Promise<void> => {
  await tx
    .update(sessionsTable)
    .set({ expiresAt: newExpiry })
    .where(eq(sessionsTable.id, sessionID))
}
