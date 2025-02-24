import { z } from "zod";
import { artistSchema } from "./artist";

export const concertSchema = z.object({
	id: z.number().positive(),
	from: z.coerce.date(),
	to: z.coerce.date(),
	artist: artistSchema,
})

export const updateConcertSchema = z.object({
	artistID: z.number().int().positive(),
	from: z.coerce.date(),
	to: z.coerce.date()
})

export type Concert = z.infer<typeof concertSchema>
export type UpdateConcert = z.infer<typeof updateConcertSchema>

/***
 * @description Returns the earliest concert within an input concerts array.
 * If the input concerts array is empty, null is returned.
 */
export const earliestConcert = (concerts: Concert[]): Concert | null => {
	if (concerts.length <= 0) return null

	return concerts.reduce((prev, curr) =>
		curr.from.getDate() < prev.from.getDate() ? curr : prev
	)
}

/***
 * @description Returns the latest concert within an input concerts array.
 * If the input concerts array is empty, null is returned.
 */
export const latestConcert = (concerts: Concert[]): Concert | null => {
	if (concerts.length <= 0) return null

	return concerts.reduce((prev, curr) =>
		curr.from.getDate() > prev.from.getDate() ? curr : prev
	)
}
