import { db } from "@/shared/db/db";
import type { CreateSessionDTO, Session } from "./session.dto";
import type { SessionRepository } from "./session.repository";
import { sessionsTable } from "@/shared/db/schema/session";
import { eq } from "drizzle-orm";

export class SQLiteSessionRepository implements SessionRepository {
  insert = async (session: CreateSessionDTO): Promise<void> => {
    await db
      .insert(sessionsTable)
      .values(session)
  }

  getByID = async (sessionID: string): Promise<Session | null> => {
    const results = await db
      .select()
      .from(sessionsTable)
      .where(eq(sessionsTable.id, sessionID))

    if (results.length <= 0) {
      return null
    }

    return results[0]
  }

  delete = async (sessionID: string): Promise<void> => {
    await db.delete(sessionsTable).where(eq(sessionsTable.id, sessionID))
  }

  setExpiry = async (sessionID: string, newExpiry: Date): Promise<void> => {
    await db
      .update(sessionsTable)
      .set({ expiresAt: newExpiry })
      .where(eq(sessionsTable.id, sessionID))
  }
}
