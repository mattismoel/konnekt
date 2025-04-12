import { z } from "zod";
import { createListResult, type ListResult } from "./list-result";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { createUrl, requestAndParse, type Query } from "./api";

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
 */
export const fetchVenues = async (fetchFn: typeof fetch, query?: Query): Promise<ListResult<Venue>> => {
	const venues = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/venues`, query),
		createListResult(venueSchema),
		"Could not fetch venues"
	)

	return venues
}

export const createVenue = async (fetchFn: typeof fetch, form: z.infer<typeof venueForm>) => {
	const venue = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/venues`),
		venueSchema,
		"Could not create venue",
		{ bodySchema: venueForm, body: form },
		"POST"
	)

	return venue
}

export const deleteVenue = async (fetchFn: typeof fetch, id: number) => {
	await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/venues/${id}`),
		undefined,
		"Could not delete venue",
		undefined,
		"DELETE"
	)
}

export const editVenue = async (fetchFn: typeof fetch, id: number, form: z.infer<typeof venueForm>) => {
	const venue = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/venues/${id}`),
		venueSchema,
		"Could not edit venue",
		{ bodySchema: venueForm, body: form },
		"PUT"
	)

	return venue
}
