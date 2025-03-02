import { z } from "zod";
import { createListResult, type ListResult } from "./list-result";
import { PUBLIC_BACKEND_URL } from "$env/static/public";

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
	if (!res.ok) throw new Error("Could not list venues")

	const result = createListResult(venueSchema).parse(await res.json())

	return result
}
