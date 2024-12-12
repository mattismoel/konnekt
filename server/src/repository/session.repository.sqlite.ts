import { db } from "@/shared/db/db";
import type { CreateSessionDTO, Session } from "@/dto/session.dto";
import type { SessionRepository } from "./session.repository";
import { deleteSessionTx, getSessionByIDTx, insertSessionTx, setSessionExpiryTx } from "@/shared/db/session";

export const createSQLiteSessionRepository = (): SessionRepository => {
  const getSessionByID = async (sessionID: string): Promise<Session | null> => {
    return await db.transaction(async (tx) => {
      return await getSessionByIDTx(tx, sessionID)
    })
  }

  const insertSession = async (session: CreateSessionDTO): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await insertSessionTx(tx, session)
    })
  }

  const deleteSession = async (sessionID: string): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await deleteSessionTx(tx, sessionID)
    })
  }

  const setSessionExpiry = async (sessionID: string, newExpiry: Date): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await setSessionExpiryTx(tx, sessionID, newExpiry)
    })
  }

  return {
    insertSession,
    getSessionByID,
    deleteSession,
    setSessionExpiry
  }
}
