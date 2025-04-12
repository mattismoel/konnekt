import { z } from "zod";
import { concertForm, concertSchema } from "./concert";
import { venueSchema } from "./venue";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { createListResult, type ListResult } from "./list-result";
import { APIError, apiErrorSchema } from "./error";
import { startOfToday } from "date-fns";
import { createUrl, requestAndParse, type Query } from "./api";

export const eventSchema = z.object({
	id: z.number().positive(),
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.string().url(),
	imageUrl: z.string().optional().or(z.string().url().optional()),
	concerts: concertSchema.array(),
	venue: venueSchema
})

export const eventForm = z.object({
	title: z.string().nonempty({ message: "Eventtitel skal være defineret" }),
	description: z.string().nonempty({ message: "Eventbeskrivelse skal være defineret" }),
	ticketUrl: z.string().nonempty({ message: "Billet-URL skal være defineret" }),
	image: z.instanceof(File).nullable(),
	venueId: z.number().positive(),
	concerts: concertForm.array().min(1, { message: "Et event skal have mindst én koncert" })
});

const createEventSchema = eventForm
	.omit({
		image: true
	})
	.extend({
		imageUrl: z.string().url()
	})

const updateEventSchema = createEventSchema

export type Event = z.infer<typeof eventSchema>

export const createEvent = async (fetchFn: typeof fetch, form: z.infer<typeof eventForm>): Promise<Event> => {
	const { image, concerts, ...rest } = form

	if (!image) throw new APIError(400, "Could not create event", "Cover image must be set")

	const imageUrl = await uploadEventCoverImage(image)

	const event = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}`),
		eventSchema,
		"Could not create event",
		{
			bodySchema: createEventSchema,
			body: { ...rest, concerts, imageUrl }
		},
	)

	return event
}

export const updateEvent = async (
	fetchFn: typeof fetch,
	form: z.infer<typeof eventForm>,
	eventId: number,
): Promise<Event> => {
	const { data, success, error } = eventForm.safeParse(form)
	if (!success) throw error

	const imageUrl = data.image ? await uploadEventCoverImage(data.image) : undefined
	const { image, concerts, ...rest } = data;

	const event = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/events/${eventId}`),
		eventSchema,
		"Could not update event",
		{
			bodySchema: updateEventSchema,
			body: { ...rest, concerts, imageUrl }
		},
		"PUT"
	)

	return event
}

export const uploadEventCoverImage = async (file: File, init?: RequestInit): Promise<string> => {
	const formData = new FormData()

	formData.append("image", file)

	const res = await fetch(`${PUBLIC_BACKEND_URL}/events/image`, {
		...init,
		method: "POST",
		credentials: "include",
		body: formData,
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not upload event cover image", err.message)
	}

	const url = await res.text()

	return url
}

export const listEvents = async (fetchFn: typeof fetch, query: Query): Promise<ListResult<Event>> => {
	const result = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/events`, query),
		createListResult(eventSchema),
		"Could not fetch events"
	)

	return result
}

/**
 * @description Returns all upcoming events, if any.
 */
export const listUpcomingEvents = async (fetchFn: typeof fetch): Promise<ListResult<Event>> => {
	const result = await listEvents(fetchFn, {
		filter: ["from_date" + ">=" + startOfToday().toISOString()]
	})

	return result
}

/**
 * @description Returns all upcoming events, if any.
 */
export const listPreviousEvents = async (fetchFn: typeof fetch): Promise<ListResult<Event>> => {
	const result = await listEvents(fetchFn, {
		filter: ["from_date" + "<" + startOfToday().toISOString()]
	})

	return result
}


export const eventById = async (fetchFn: typeof fetch, id: number): Promise<Event> => {
	const event = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/events/${id}`),
		eventSchema,
	)

	return event
}

export const artistEvents = async (fetchFn: typeof fetch, artistId: number): Promise<ListResult<Event>> => {
	const result = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/events`, {
			filter: [
				"from_date" + ">=" + startOfToday().toISOString(),
				"artist_id" + "=" + artistId.toString()
			]
		}),
		createListResult(eventSchema),
	)

	return result
}

export const deleteEvent = async (fetchFn: typeof fetch, id: number) => {
	await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/events/${id}`),
		undefined,
		"Could not delete event",
		undefined,
		"DELETE"
	)
}
