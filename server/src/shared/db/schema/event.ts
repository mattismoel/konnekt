import { integer, primaryKey, sqliteTable, text } from "drizzle-orm/sqlite-core";
import { addressesTable } from "./address";
import { genresTable } from "./genre";

export const eventsTable = sqliteTable("event", {
  id: integer().primaryKey({ autoIncrement: true }),
  title: text().notNull(),
  description: text().notNull(),
  coverImageUrl: text(),
  fromDate: integer({ mode: "timestamp" }).notNull(),
  toDate: integer({ mode: "timestamp" }).notNull(),
  addressID: integer().notNull().references(() => addressesTable.id)
})

export const eventsGenresTable = sqliteTable("events_genres", {
  eventID: integer().notNull().references(() => eventsTable.id),
  genreID: integer().notNull().references(() => genresTable.id)
}, (table) => ({
  pk: primaryKey({ columns: [table.eventID, table.genreID] })
}))
