import { z } from "zod";
import { artistSchema } from "./artist";

export const concertSchema = z.object({
	id: z.number().positive(),
	from: z.coerce.date(),
	to: z.coerce.date(),
	artist: artistSchema,
})

export type Concert = z.infer<typeof concertSchema>
