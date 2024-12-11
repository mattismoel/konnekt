import { db } from "@/shared/db/db";
import type { CreateSessionDTO, Session } from "@/dto/session.dto";
import type { SessionRepository } from "./session.repository";
import { deleteSessionTx, getSessionByIDTx, insertSessionTx, setSessionExpiryTx } from "@/shared/db/session";

export class SQLiteSessionRepository implements SessionRepository {
  getByID = async (sessionID: string): Promise<Session | null> => {
    return await db.transaction(async (tx) => {
      return await getSessionByIDTx(tx, sessionID)
    })
  }

  insert = async (session: CreateSessionDTO): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await insertSessionTx(tx, session)
    })
  }

  delete = async (sessionID: string): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await deleteSessionTx(tx, sessionID)
    })
  }

  setExpiry = async (sessionID: string, newExpiry: Date): Promise<void> => {
    return db.transaction(async (tx) => {
      return await setSessionExpiryTx(tx, sessionID, newExpiry)
    })
  }
}
