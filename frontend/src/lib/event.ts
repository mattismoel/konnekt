import { z } from "zod";
import { concertForm, concertSchema } from "./concert";
import { venueSchema } from "./venue";

export const eventSchema = z.object({
	id: z.number().positive(),
	title: z.string(),
	description: z.string(),
	coverImageUrl: z.string().url(),
	concerts: concertSchema.array(),
	venue: venueSchema
})

export const eventForm = z.object({
	title: z.string(),
	description: z.string(),
	coverImageUrl: z.string().url(),
	venueId: z.number().positive(),
	concerts: concertForm.array().min(1)
});

export type Event = z.infer<typeof eventSchema>
