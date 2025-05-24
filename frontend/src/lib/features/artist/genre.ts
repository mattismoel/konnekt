import { idSchema, requestAndParse } from "@/lib/api"
import { createListResult, type ListResult } from "@/lib/query"
import { createUrl, type Query } from "@/lib/url"
import { z } from "zod"
import type { Event } from "../event/event"
import { eventsArtists } from "./events-artists"

export const genreSchema = z.object({
	id: idSchema,
	name: z.string()
})

export type Genre = z.infer<typeof genreSchema>

const createGenreSchema = z.object({
	name: z.string().nonempty()
})

export const listGenres = async (query?: Query): Promise<ListResult<Genre>> => {
	const result = requestAndParse(
		createUrl("/api/genres", query),
		createListResult(genreSchema)
	)

	return result
}

export const createGenre = async (name: string): Promise<void> => {
	await requestAndParse(
		createUrl("/api/genres"),
		undefined,
		"Could not create genre",
		{ bodySchema: createGenreSchema, body: { name } },
		"POST",
	)
}

/**
	* @author https://stackoverflow.com/questions/2218999/how-to-remove-all-duplicates-from-an-array-of-objects
	*/
export const eventGenres = (e: Event) => {
	const artists = e.concerts.map(c => c.artist)

	const genres = artists
		.flatMap((artist) => artist.genres)
		.filter((value, index, self) => index === self.findIndex((t) => t.id === value.id))
		.sort((a, b) => {
			const nameA = a.name.toUpperCase();
			const nameB = b.name.toUpperCase();

			return nameA < nameB ? -1 : nameA > nameB ? 1 : 0;
		})

	return genres
}
