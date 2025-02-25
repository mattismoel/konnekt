import { z } from "zod";
import { concertSchema } from "./concert";
import { venueSchema } from "./venue";

export const eventSchema = z.object({
	id: z.number().positive(),
	title: z.string(),
	description: z.string(),
	coverImageUrl: z.string().url(),
	concerts: concertSchema.array(),
	venue: venueSchema
})

export type Event = z.infer<typeof eventSchema>
