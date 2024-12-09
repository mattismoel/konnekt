import { db } from "@/shared/db/db";
import type { GenreRepository } from "./genre.repository";
import { genresTable } from "@/shared/db/schema/genre";
import { eq } from "drizzle-orm";
import { eventsGenresTable } from "@/shared/db/schema/event";

export class SQLiteGenreRepository implements GenreRepository {
  insert = async (genre: string): Promise<void> => {
    await db
      .insert(genresTable)
      .values({ name: genre })
      .onConflictDoNothing()
  }

  getAll = async (): Promise<string[]> => {
    const genres = await db
      .select({ name: genresTable.name })
      .from(genresTable)

    return genres.map(genre => genre.name)
  }

  delete = async (genre: string): Promise<void> => {
    await db
      .delete(genresTable)
      .where(eq(genresTable.name, genre))
  }

  eventGenres = async (eventID: number): Promise<string[]> => {
    const results = await db.
      select()
      .from(genresTable)
      .innerJoin(eventsGenresTable, eq(genresTable.id, eventsGenresTable.genreID))
      .where(eq(eventsGenresTable.eventID, eventID))

    const genres = results.map(res => res.genre.name)

    return genres
  }
}
