import { db } from "@/shared/db/db";
import type { GenreRepository } from "./genre.repository";
import { deleteGenreTx, getAllGenresTx, getEventGenresTx, insertGenreTx } from "@/shared/db/genre";
import type { GenreDTO } from "@/dto/genre.dto";

export const createSQLiteGenreRepository = (): GenreRepository => {
  const insertGenre = async (genre: string): Promise<GenreDTO> => {
    return await db.transaction(async (tx) => await insertGenreTx(tx, genre))
  }

  const listGenres = async (): Promise<GenreDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getAllGenresTx(tx)
    })
  }

  const deleteGenre = async (genre: string): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await deleteGenreTx(tx, genre)
    })
  }

  const getEventGenres = async (eventID: number): Promise<GenreDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getEventGenresTx(tx, eventID)
    })
  }

  return {
    insertGenre,
    listGenres,
    getEventGenres,
    deleteGenre,
  }
}
