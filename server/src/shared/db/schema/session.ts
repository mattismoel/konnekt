import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";
import { usersTable } from "./user";

export const sessionsTable = sqliteTable("session", {
  id: text("id").primaryKey(),
  userID: integer("user_id").notNull().references(() => usersTable.id),
  expiresAt: integer("expires_at", { mode: "timestamp" }).notNull()
})
