import { z } from "zod";

export const venueSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	city: z.string(),
	countryCode: z.string(),
})

export type Venue = z.infer<typeof venueSchema>
