import { z } from "zod";

export const artistSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	imageUrl: z.string().url(),
	description: z.string(),
	genres: z.string().array(),
	socials: z.string().array()
})

export type Artist = z.infer<typeof artistSchema>
