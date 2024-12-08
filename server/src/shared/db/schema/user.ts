import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const usersTable = sqliteTable("user", {
  id: integer().primaryKey({ autoIncrement: true }),
  email: text().notNull().unique(),
  firstName: text().notNull(),
  lastName: text().notNull()
})
