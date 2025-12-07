import { baseFields } from "$lib/model";
import { z } from "zod";

export const genreSchema = z.object({
	...baseFields.shape,
	name: z.string().nonempty()
});

export const socialSchema = z.object({
	...baseFields.shape,
	url: z.url()
});

export const artistSchema = z.object({
	...baseFields.shape,
	name: z.string().nonempty(),
	description: z.string().nonempty(),
	cover: z.url(),
	previewUrl: z.union([z.url(), z.literal("")]),
	genres: genreSchema.array().min(1),
	socials: socialSchema.array().optional()
});

export type Artist = z.infer<typeof artistSchema>;
export type Genre = z.infer<typeof genreSchema>;
export type Social = z.infer<typeof socialSchema>;
