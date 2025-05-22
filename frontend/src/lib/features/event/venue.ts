import { idSchema, requestAndParse, type ID } from "@/lib/api";
import { createListResult, type ListResult } from "@/lib/query";
import { createUrl, type Query } from "@/lib/url";
import { z } from "zod";

export const venueSchema = z.object({
	id: idSchema,
	name: z.string(),
	city: z.string(),
	countryCode: z.string(),
})

export type Venue = z.infer<typeof venueSchema>

export const venueForm = z.object({
	name: z.string().nonempty(),
	city: z.string().nonempty(),
	countryCode: z.string().nonempty(),
})

export type VenueFormValues = z.infer<typeof venueForm>


/**
 * @description Lists venues.
 */
export const listVenues = async (query?: Query): Promise<ListResult<Venue>> => {
	const venues = await requestAndParse(
		createUrl("/api/venues", query),
		createListResult(venueSchema),
		"Could not fetch venues"
	)

	return venues
}

export const createVenue = async (form: VenueFormValues) => {
	const venue = await requestAndParse(
		createUrl("/api/venues"),
		venueSchema,
		"Could not create venue",
		{ bodySchema: venueForm, body: form },
		"POST"
	)

	return venue
}

export const deleteVenue = async (id: ID) => {
	await requestAndParse(
		createUrl(`/api/venues/${id}`),
		undefined,
		"Could not delete venue",
		undefined,
		"DELETE"
	)
}

export const editVenue = async (id: ID, form: VenueFormValues) => {
	const venue = requestAndParse(
		createUrl(`/api/venues/${id}`),
		venueSchema,
		"Could not edit venue",
		{ bodySchema: venueForm, body: form },
		"PUT"
	)

	return venue
}

export const venueById = async (venueId: ID) => {
	const venue = await requestAndParse(
		createUrl(`/api/venues/${venueId}`),
		venueSchema,
		"Could not get venue by ID"
	)

	return venue
}
