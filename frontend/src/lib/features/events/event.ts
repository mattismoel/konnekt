import z from "zod";
import { artistSchema } from "../artists/artist";
import { baseFields } from "$lib/model";
import { venueSchema } from "../venues/venue";

export const concertSchema = z.object({
	...baseFields.shape,
	fromDate: z.coerce.date(),
	toDate: z.coerce.date(),
	artist: artistSchema
});

export const eventSchema = z.object({
	...baseFields.shape,
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.url(),
	cover: z.url(),
	venue: venueSchema,
	concerts: concertSchema.array(),
	isPublic: z.boolean(),
});

export type Event = z.infer<typeof eventSchema>;
export type Concert = z.infer<typeof concertSchema>;

/***
 * @description Returns the earliest concert within an input concerts array.
 * If the input concerts array is empty, null is returned.
 */
export const earliestConcert = (concerts: Concert[]): Concert | null => {
	if (concerts.length === 0) return null;

	return concerts.reduce((prev, curr) =>
		curr.fromDate.getDate() < prev.fromDate.getDate() ? curr : prev
	);
};

export const eventGenres = (event: Event) => {
	const genres = event.concerts
		.flatMap((concert) => concert.artist.genres)
		.filter((genre, i, self) => i === self.findIndex((x) => x.id === genre.id))
		.sort((a, b) => {
			const nameA = a.name.toUpperCase();
			const nameB = b.name.toUpperCase();

			return nameA < nameB ? -1 : nameA > nameB ? 1 : 0;
		});

	return genres;
};
