import { z } from "zod";
import { genreSchema } from "./genre";

export const artistSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	imageUrl: z.string().url(),
	description: z.string(),
	genres: genreSchema.array(),
	socials: z.string().url().array()
})

export type Artist = z.infer<typeof artistSchema>

export const artistFormSchema = z.object({
	name: z.string(),
	description: z.string(),
	imageUrl: z.string().url(),
	genreIds: z.number().positive().array(),
	socials: z.string().url().array(),
})

export type ArtistForm = z.infer<typeof artistFormSchema>
