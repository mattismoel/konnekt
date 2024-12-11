import { db } from "@/shared/db/db";
import type { GenreRepository } from "./genre.repository";
import { deleteGenreTx, getAllGenresTx, getEventGenresTx, insertGenreTx } from "@/shared/db/genre";
import type { GenreDTO } from "@/dto/genre.dto";

export class SQLiteGenreRepository implements GenreRepository {
  insert = async (genre: string): Promise<GenreDTO> => {
    return await db.transaction(async (tx) => await insertGenreTx(tx, genre))
  }

  getAll = async (): Promise<GenreDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getAllGenresTx(tx)
    })
  }

  delete = async (genre: string): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await deleteGenreTx(tx, genre)
    })
  }

  eventGenres = async (eventID: number): Promise<GenreDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getEventGenresTx(tx, eventID)
    })
  }
}
