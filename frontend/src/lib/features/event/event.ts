import { z } from "zod";
import { concertForm, concertSchema } from "./concert";
import { venueSchema } from "./venue";
import { APIError, apiErrorSchema, idSchema, requestAndParse, type ID } from "@/lib/api";
import { createUrl, type Query } from "@/lib/url";
import { createListResult, type ListResult } from "@/lib/query";
import { startOfToday } from "date-fns";
import { env } from "../../env";

export const eventSchema = z.object({
	id: idSchema,
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.string().url(),
	imageUrl: z.string().optional().or(z.string().url().optional()),
	concerts: concertSchema.array(),
	venue: venueSchema
})

export type Event = z.infer<typeof eventSchema>

export const eventForm = z.object({
	title: z.string().nonempty({ message: "Eventtitel skal være defineret" }),
	description: z.string().nonempty({ message: "Eventbeskrivelse skal være defineret" }),
	ticketUrl: z.string().nonempty({ message: "Billet-URL skal være defineret" }),
	venueId: idSchema,
	concerts: concertForm.array().min(1, { message: "Et event skal have mindst én koncert" }),
	image: z.instanceof(File).optional()
});

export type EventFormValues = z.infer<typeof eventForm>

const createEventSchema = eventForm
	.omit({ image: true })
	.extend({ imageUrl: z.string().url() })

export const createEvent = async (form: EventFormValues): Promise<Event> => {
	const { data: formData, error: formError } = eventForm.safeParse(form)
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

const updateEventSchema = eventForm
	.omit({ image: true })
	.extend({ imageUrl: z.string().url().optional() })

export const updateEvent = async (form: EventFormValues, eventId: ID): Promise<Event> => {
	const { data, success, error } = eventForm.safeParse(form)
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
		createUrl(`/api/events`, query),
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


export const eventById = async (id: ID): Promise<Event> => {
	const event = await requestAndParse(
		createUrl(`/api/events/${id}`),
		eventSchema,
	)

	return event
}

export const artistEvents = async (artistId: ID): Promise<ListResult<Event>> => {
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

export const deleteEvent = async (id: ID) => {
	await requestAndParse(
		createUrl(`/api/events/${id}`),
		undefined,
		"Could not delete event",
		undefined,
		"DELETE"
	)
}
