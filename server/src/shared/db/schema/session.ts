import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";
import { usersTable } from "./user";

export const sessionsTable = sqliteTable("session", {
  id: text().primaryKey(),
  userID: integer().notNull().references(() => usersTable.id),
  expiresAt: integer({ mode: "timestamp" }).notNull()
})
