import type { GenreDTO } from "@/dto/genre.dto"

export interface GenreRepository {
  listGenres(): Promise<GenreDTO[]>
  insertGenre(genre: string): Promise<GenreDTO>
  deleteGenre(genre: string): Promise<void>
  getEventGenres(genreID: number): Promise<GenreDTO[]>
}
