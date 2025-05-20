import { idSchema, requestAndParse } from "@/lib/api"
import { createListResult, type ListResult } from "@/lib/query"
import { createUrl, type Query } from "@/lib/url"
import { z } from "zod"

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

