import type { RequestHandler } from "express";
import type { GenreService } from "@/service/genre.service";

export class GenreController {
  constructor(
    private readonly genreService: GenreService
  ) { }

  getAll: RequestHandler = async (req, res, next) => {
    const genres = await this.genreService.getAll()

    res.json(genres)
  }
}
