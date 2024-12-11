import type { GenreDTO } from "@/dto/genre.dto"

export interface GenreRepository {
  getAll(): Promise<GenreDTO[]>
  insert(genre: string): Promise<GenreDTO>
  delete(genre: string): Promise<void>
  eventGenres(genreID: number): Promise<GenreDTO[]>
}
