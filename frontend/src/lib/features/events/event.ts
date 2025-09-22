import { z } from "zod"
import { artistSchema } from "../artists/artist"
import { venueSchema } from "../venues/venue"

export const concertSchema = z.object({
	id: z.int().positive(),
	artist: artistSchema,
	from: z.coerce.date(),
	to: z.coerce.date(),
})

export const eventSchema = z.object({
	id: z.int().positive(),
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.url(),
	imageUrl: z.url(),
	concerts: concertSchema.array(),
	venue: venueSchema,
})


export type Event = z.infer<typeof eventSchema>
export type Concert = z.infer<typeof concertSchema>
