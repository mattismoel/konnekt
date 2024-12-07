import { integer, sqliteTable, text } from "drizzle-orm/sqlite-core";

export const addressesTable = sqliteTable("address", {
  id: integer().primaryKey({ autoIncrement: true }),
  country: text().notNull(),
  city: text().notNull(),
  street: text().notNull(),
  houseNumber: text().notNull()
})
