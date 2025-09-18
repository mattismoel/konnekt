import { z } from "zod"
import { artist } from "../artists/artist"
import { venue } from "../venues/venue"

export const concert = z.object({
	id: z.int().positive(),
	artist: artist,
	from: z.coerce.date(),
	to: z.coerce.date(),
})

export const event = z.object({
	id: z.int().positive(),
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.url(),
	imageUrl: z.url(),
	concerts: concert.array(),
	venue: venue,
})


export type Event = z.infer<typeof event>
