import { z } from "zod";
import { concertForm, concertSchema } from "./concert";
import { venueSchema } from "./venue";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { createListResult, type ListResult } from "./list-result";

export const eventSchema = z.object({
	id: z.number().positive(),
	title: z.string(),
	description: z.string(),
	coverImageUrl: z.string().url(),
	concerts: concertSchema.array(),
	venue: venueSchema
})

export const eventForm = z.object({
	title: z.string(),
	description: z.string(),
	coverImageUrl: z.string().url(),
	venueId: z.number().positive(),
	concerts: concertForm.array().min(1)
});

export type Event = z.infer<typeof eventSchema>


export const listEvents = async (params: URLSearchParams): Promise<ListResult<Event>> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/events?` + params.toString())
	if (!res.ok) {
		throw new Error("Could not list events")
	}

	const result = createListResult(eventSchema).parse(await res.json())

	return result
}

export const eventById = async (id: number): Promise<Event> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/events/${id}`)

	if (!res.ok) throw new Error(`Could not get event with id ${id}`)

	const event = eventSchema.parse(await res.json())

	return event
}
