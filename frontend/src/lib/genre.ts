import { z } from "zod";
import { createListResult, type ListResult } from "./list-result";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { requestAndParse } from "./api";
import { createUrl, type Query } from "./url";

export const genreSchema = z.object({
	id: z.number().positive(),
	name: z.string()
})

export type Genre = z.infer<typeof genreSchema>

const createGenreSchema = z.object({
	name: z.string().nonempty()
})

export const listGenres = async (fetchFn: typeof fetch, query?: Query): Promise<ListResult<Genre>> => {
	const result = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/genres`, query),
		createListResult(genreSchema)
	)

	return result
}

export const createGenre = async (fetchFn: typeof fetch, name: string): Promise<void> => {
	await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/genres`),
		undefined,
		"Could not create genre",
		{ bodySchema: createGenreSchema, body: { name } },
		"POST",
	)
}
