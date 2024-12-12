import { eq, inArray } from "drizzle-orm"
import type { TX } from "./db"
import { genresTable } from "./schema/genre"
import { eventsGenresTable } from "./schema/event"
import type { GenreDTO } from "@/dto/genre.dto"

export const insertGenreTx = async (tx: TX, genre: string): Promise<GenreDTO> => {
  const insertedGenre = await tx
    .insert(genresTable)
    .values({ name: genre })
    .onConflictDoNothing()
    .returning()

  return insertedGenre[0]
}

export const getAllGenresTx = async (tx: TX): Promise<GenreDTO[]> => {
  const genres = await tx
    .select()
    .from(genresTable)

  return genres
}

export const deleteGenreTx = async (tx: TX, genre: string): Promise<void> => {
  await tx
    .delete(genresTable)
    .where(eq(genresTable.name, genre))
}

/**
 * @description Inserts genres into the database, if they do not already exist.
 * Only inserted genres are returned.
 */
export const insertGenresTx = async (tx: TX, genreNames: string[]): Promise<GenreDTO[]> => {
  const genres = await Promise.all(
    genreNames.map(async genreName => await insertGenreTx(tx, genreName))
  )

  return genres
}

/**
 * @description Gets all genres related to an event.
 */
export const getEventGenresTx = async (tx: TX, eventID: number): Promise<GenreDTO[]> => {
  const results = await tx.
    select()
    .from(genresTable)
    .innerJoin(eventsGenresTable, eq(genresTable.id, eventsGenresTable.genreID))
    .where(eq(eventsGenresTable.eventID, eventID))

  const genres = results.map(res => res.genre)

  return genres
}

/**
 * @description Relates the input genres to a given event.
 */
export const relateGenresToEventTx = async (tx: TX, eventID: number, genreNames: string[]): Promise<void> => {
  const genres = await tx
    .select({ id: genresTable.id })
    .from(genresTable)
    .where(inArray(genresTable.name, genreNames))

  await tx
    .insert(eventsGenresTable)
    .values(genres.map(genre => ({ eventID, genreID: genre.id })))
}
