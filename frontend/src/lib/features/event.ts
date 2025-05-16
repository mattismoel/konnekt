import { z } from "zod";
import { concertForm, concertSchema } from "./concert";
import { venueSchema } from "./venue";
import { APIError, apiErrorSchema, requestAndParse } from "@/lib/api";
import { createUrl, type Query } from "@/lib/url";
import { createListResult, type ListResult } from "@/lib/query";
import { startOfToday } from "date-fns";
import { env } from "../env";
// import { PUBLIC_PROXY_ORIGIN } from "$env/static/public";

export const eventSchema = z.object({
	id: z.number().positive(),
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.string().url(),
	imageUrl: z.string().optional().or(z.string().url().optional()),
	concerts: concertSchema.array(),
	venue: venueSchema
})

const eventForm = z.object({
	title: z.string().nonempty({ message: "Eventtitel skal være defineret" }),
	description: z.string().nonempty({ message: "Eventbeskrivelse skal være defineret" }),
	ticketUrl: z.string().nonempty({ message: "Billet-URL skal være defineret" }),
	venueId: z.number().positive(),
	concerts: concertForm.array().min(1, { message: "Et event skal have mindst én koncert" })
});

export const createEventForm = eventForm
	.extend({ image: z.instanceof(File).optional() })

export const editEventForm = eventForm
	.extend({ image: z.instanceof(File).optional() })


const createEventSchema = createEventForm
	.omit({ image: true })
	.extend({ imageUrl: z.string().url() })

const updateEventSchema = editEventForm
	.omit({ image: true })
	.extend({ imageUrl: z.string().url().optional() })

export type Event = z.infer<typeof eventSchema>


export const createEvent = async (form: z.infer<typeof createEventForm>,): Promise<Event> => {
	const { data: formData, error: formError } = createEventForm.safeParse(form)
	if (formError) throw formError

	const { image, ...rest } = formData
	if (!image) throw new APIError(400, "Could not create event", "Cover image must be set")

	const imageUrl = await uploadEventCoverImage(image)

	const event = await requestAndParse(
		createUrl(`/api/events`),
		eventSchema,
		"Could not create event",
		{ body: { ...rest, imageUrl }, bodySchema: createEventSchema },
		"POST",
	)

	return event
}

export const updateEvent = async (
	form: z.infer<typeof editEventForm>,
	eventId: number,
): Promise<Event> => {
	const { data, success, error } = editEventForm.safeParse(form)
	if (!success) throw error

	const imageUrl = data.image ? await uploadEventCoverImage(data.image) : undefined
	const { image, ...rest } = data;

	const event = await requestAndParse(
		createUrl(`/api/events/${eventId}`),
		eventSchema,
		"Could not update event",
		{
			bodySchema: updateEventSchema,
			body: { ...rest, imageUrl }
		},
		"PUT"
	)

	return event
}

export const uploadEventCoverImage = async (file: File, init?: RequestInit): Promise<string> => {
	const formData = new FormData()

	formData.append("image", file)

	const res = await fetch(`/api/events/image`, {
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

export const listEvents = async (query: Query,): Promise<ListResult<Event>> => {
	const result = requestAndParse(
		createUrl(`${env.VITE_SERVER_ORIGIN}/api/events`, query),
		createListResult(eventSchema),
		"Could not fetch events"
	)

	return result
}

export const listUpcomingEvents = async (): Promise<ListResult<Event>> => {
	return listEvents({ filter: ["from_date" + ">=" + startOfToday().toISOString()] })
}



// /**
//  * @description Returns all upcoming events, if any.
//  */
// export const listUpcomingEvents = async (): Promise<ListResult<Event>> => {
// 	const result = await listEvents(
// 		fetchFn,
// 		{
// 			filter: ["from_date" + ">=" + startOfToday().toISOString()]
// 		},
// 	)
//
// 	return result
// }

/**
 * @description Returns all upcoming events, if any.
 */
export const listPreviousEvents = async (): Promise<ListResult<Event>> => {
	const result = await listEvents({
		filter: ["from_date" + "<" + startOfToday().toISOString()]
	})

	return result
}


export const eventById = async (id: number): Promise<Event> => {
	const event = await requestAndParse(
		createUrl(`${""}/api/events/${id}`),
		eventSchema,
	)

	return event
}

export const artistEvents = async (artistId: number): Promise<ListResult<Event>> => {
	const result = requestAndParse(
		createUrl(`/api/events`, {
			filter: [
				"from_date" + ">=" + startOfToday().toISOString(),
				"artist_id" + "=" + artistId.toString()
			]
		}),
		createListResult(eventSchema),
	)

	return result
}

export const deleteEvent = async (id: number,) => {
	await requestAndParse(
		createUrl(`/events/${id}`),
		undefined,
		"Could not delete event",
		undefined,
		"DELETE"
	)
}
