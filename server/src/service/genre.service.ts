import type { GenreRepository } from "@/repository/genre.repository";

export type GenreService = {
  listGenres(): Promise<string[]>
}

export const createGenreService = (genreRepo: GenreRepository): GenreService => {
  const listGenres = async (): Promise<string[]> => {
    const genres = await genreRepo.listGenres()
    return genres.map(genre => genre.name)
  }

  return {
    listGenres,
  }
}
