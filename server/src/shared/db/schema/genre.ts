import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const genresTable = sqliteTable("genre", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  name: text("name").notNull().unique()
})
