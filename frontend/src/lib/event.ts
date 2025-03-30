import { z } from "zod";
import { concertForm, concertSchema } from "./concert";
import { venueSchema } from "./venue";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { createListResult, type ListResult } from "./list-result";
import { APIError, apiErrorSchema } from "./error";
import { startOfToday } from "date-fns";

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

export type Event = z.infer<typeof eventSchema>

export const createEvent = async (form: z.infer<typeof eventForm>, init?: RequestInit): Promise<Event> => {
	const { image, concerts, ...rest } = form

	if (!image) throw new APIError(400, "Could not create event", "Cover image must be set")

	const imageUrl = await uploadEventCoverImage(image)

	let res = await fetch(`${PUBLIC_BACKEND_URL}/events`, {
		...init,
		method: "POST",
		credentials: "include",
		body: JSON.stringify({
			...rest,
			concerts: concerts.map(c => ({
				...c,
				from: c.from.toISOString(),
				to: c.to.toISOString(),
			})),
			imageUrl
		}),
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not update event", err.message)
	}

	const event = eventSchema.parse(await res.json())

	return event
}

export const updateEvent = async (form: z.infer<typeof eventForm>, eventId: number, init?: RequestInit): Promise<Event> => {
	const { data, success, error } = eventForm.safeParse(form)
	if (!success) throw error

	const imageUrl = data.image ? await uploadEventCoverImage(data.image) : undefined
	const { image, concerts, ...rest } = data;

	const res = await fetch(`${PUBLIC_BACKEND_URL}/events/${eventId}`, {
		...init,
		method: "PUT",
		credentials: "include",
		body: JSON.stringify({
			...rest,
			concerts: concerts.map(c => ({
				...c,
				from: c.from.toISOString(),
				to: c.to.toISOString(),
			})),
			imageUrl
		})
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not update event", err.message)
	}

	const event = eventSchema.parse(await res.json())
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

export const listEvents = async (params: URLSearchParams): Promise<ListResult<Event>> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/events?` + params.toString())
	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not list events", err.message)
	}

	const result = createListResult(eventSchema).parse(await res.json())

	return result
}

/**
 * @description Returns all upcoming events, if any.
 */
export const listUpcomingEvents = async (): Promise<ListResult<Event>> => {
	const result = await listEvents(new URLSearchParams({
		"filter": `from_date>=${startOfToday().toISOString()}`,
	}))

	return result
}

export const eventById = async (id: number): Promise<Event> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/events/${id}`)

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, `Could not get event`, err.message)
	}

	const event = eventSchema.parse(await res.json())

	return event
}
