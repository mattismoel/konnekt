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
	name: z
		.string()
		.nonempty({ message: "Kustnernavn skal være defineret" }),
	description: z
		.string()
		.nonempty({ message: "Kunstnerbeskreivelse skal være defineret" }),
	imageUrl: z
		.string()
		.nonempty({ message: "Kuntnerbillede skal være defineret" })
		.url({ message: "Kunstnerbillede skal være gyldigt" }),
	genreIds: z.number()
		.positive()
		.array()
		.min(1, { message: "Mindst én genre skal være valgt" }),
	socials: z
		.string()
		.nonempty()
		.url({ message: "URL skal være gyldigt" })
		.array(),
})

export type ArtistForm = z.infer<typeof artistFormSchema>
