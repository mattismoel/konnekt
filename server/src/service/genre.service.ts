import type { GenreRepository } from "@/repository/genre.repository";

export class GenreService {
  constructor(private readonly genreRepository: GenreRepository) { }

  getAll = async (): Promise<string[]> => {
    const genres = await this.genreRepository.getAll()
    return genres.map(genre => genre.name)
  }
}
