export interface GenreRepository {
  getAll(): Promise<string[]>
  insert(genre: string): Promise<void>
  delete(genre: string): Promise<void>
  eventGenres(genreID: number): Promise<string[]>
}
