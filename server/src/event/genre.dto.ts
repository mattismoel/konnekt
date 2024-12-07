import { z } from "zod";

export type GenreDTO = {
  id: number;
  name: string;
}

export const createGenreSchema = z.string().min(1)

export type CreateGenreDTO = z.infer<typeof createGenreSchema>
