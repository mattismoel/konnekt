import { z } from "zod";
import { artistSchema } from "../artist/artist";
import { idSchema } from "@/lib/api";

export const concertSchema = z.object({
	id: idSchema,
	from: z.coerce.date(),
	to: z.coerce.date(),
	artist: artistSchema,
})

export type Concert = z.infer<typeof concertSchema>

export const concertForm = z.object({
	artistID: idSchema,
	from: z.coerce.date(),
	to: z.coerce.date()
});

export const createConcertForm = concertForm
export const editConcertForm = concertForm

export type CreateConcertForm = z.infer<typeof createConcertForm>
export type EditConcertForm = z.infer<typeof editConcertForm>

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
