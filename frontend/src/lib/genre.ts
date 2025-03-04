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
	const res = await fetch(`${PUBLIC_BACKEND_URL}/genres?`, {
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
