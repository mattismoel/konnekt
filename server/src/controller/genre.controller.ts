import type { RequestHandler } from "express";
import type { GenreService } from "@/service/genre.service";

export type GenreController = {
  listGenres: RequestHandler
}

export const createGenreController = (genreService: GenreService): GenreController => {
  const listGenres: RequestHandler = async (req, res, next) => {
    const genres = await genreService.listGenres()

    res.json(genres)
  }

  return { listGenres }
}
