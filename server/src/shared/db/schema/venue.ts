import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const venuesTable = sqliteTable("venue", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  name: text("name").notNull(),
  city: text("city").notNull(),
  country: text("country").notNull(),
})
