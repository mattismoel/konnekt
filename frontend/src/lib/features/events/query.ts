import { throwAPIError } from "@/lib/api/error";
import { queryOptions } from "@tanstack/react-query";
import { eventSchema } from "./event";
import { createListResult } from "@/lib/api/list";

export const upcomingEventsQueryOpts = () => queryOptions({
	queryKey: ["events", "upcoming"],
	queryFn: async () => {
		const res = await fetch("/api/events")
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		return createListResult(eventSchema).parse(await res.json())
	}
})
export const eventByIdQueryOpts = (eventId: number) => queryOptions({
	queryKey: ["events", eventId],
	queryFn: async () => {
		const res = await fetch(`/api/events/${eventId}`)
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		return eventSchema.parse(await res.json())
	}
})

export const artistEventsQueryOpts = (artistId: number) => queryOptions({
	queryKey: ["artist-events", artistId],
	queryFn: async () => {
		const res = await fetch(`/api/artists/${artistId}/events`)
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		return createListResult(eventSchema).parse(await res.json())
	}
})
