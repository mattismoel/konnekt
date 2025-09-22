import z from "zod";

export const artistSchema = z.object({
	id: z.int().positive(),
	name: z.string().nonempty(),
	description: z.string().nonempty(),
	imageUrl: z.url(),
	previewUrl: z.url().or(z.literal("")),
	socials: z.url().array(),
	genres: z.string().array().min(1)
})

export type Artist = z.infer<typeof artistSchema>
