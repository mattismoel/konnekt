import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const genresTable = sqliteTable("genre", {
  id: integer().primaryKey({ autoIncrement: true }),
  name: text().notNull().unique()
})
