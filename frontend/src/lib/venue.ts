import { z } from "zod";
import { createListResult, type ListResult } from "./list-result";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { APIError, apiErrorSchema } from "./error";

export const venueForm = z.object({
	name: z.string().nonempty(),
	city: z.string().nonempty(),
	countryCode: z.string().nonempty(),
})

export const venueSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	city: z.string(),
	countryCode: z.string(),
})

export type Venue = z.infer<typeof venueSchema>

/**
 * @description Lists venues.
 * @param {RequestInit} init - Must be specified with {credentials: "include"} if fetching with auth cookie.
 */
export const listVenues = async (init?: RequestInit, params?: URLSearchParams): Promise<ListResult<Venue>> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/venues?` + params, init)
	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not list venues", err.message)
	}

	const result = createListResult(venueSchema).parse(await res.json())

	return result
}

export const createVenue = async (form: z.infer<typeof venueForm>, init?: RequestInit) => {
	const data = venueForm.parse(form)

	const res = await fetch(`${PUBLIC_BACKEND_URL}/venues`, {
		method: "POST",
		credentials: "include",
		body: JSON.stringify(data),
		...init
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not create venue", err.message)
	}
}

export const deleteVenue = async (id: number, init?: RequestInit) => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/venues/${id}`, {
		method: "DELETE",
		credentials: "include",
		...init,
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not create venue", err.message)
	}
}
