import { integer, primaryKey, sqliteTable, text } from "drizzle-orm/sqlite-core";
import { genresTable } from "./genre";
import { venuesTable } from "./venue";

export const eventsTable = sqliteTable("event", {
  id: integer("id").primaryKey({ autoIncrement: true }),
  title: text("title").notNull(),
  description: text("description").notNull(),
  coverImageUrl: text("cover_image_url").notNull(),
  fromDate: integer("from_date", { mode: "timestamp" }).notNull(),
  toDate: integer("to_date", { mode: "timestamp" }).notNull(),
  venueID: integer("venue_id").notNull().references(() => venuesTable.id)
})

export const eventsGenresTable = sqliteTable("events_genres", {
  eventID: integer("event_id").notNull().references(() => eventsTable.id),
  genreID: integer("genre_id").notNull().references(() => genresTable.id)
}, (table) => ({
  pk: primaryKey({ columns: [table.eventID, table.genreID] })
}))
