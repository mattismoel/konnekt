import { z } from "zod";
import { createListResult, type ListResult } from "./list-result";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { APIError, apiErrorSchema } from "./error";

export const genreSchema = z.object({
	id: z.number().positive(),
	name: z.string()
})

export type Genre = z.infer<typeof genreSchema>

export const listGenres = async (params?: URLSearchParams, init?: RequestInit): Promise<ListResult<Genre>> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/genres?` + params?.toString(), {
		...init,
		credentials: "include",
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not list genres", err.message)
	}

	const result = createListResult(genreSchema).parse(await res.json())

	return result
}

export const createGenre = async (name: string, init?: RequestInit): Promise<void> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/genres`, {
		method: "POST",
		credentials: "include",
		...init
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not create genre", err.message)
	}
}
