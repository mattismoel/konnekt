import { queryOptions } from "@tanstack/react-query";
import { listVenues, venueById } from "./venue";
import { eventById, listPreviousEvents, listUpcomingEvents } from "./event";
import type { ID } from "@/lib/api";

export const venuesQueryOpts =
	queryOptions({
		queryKey: ["venues"],
		queryFn: () => listVenues()
	})

export const upcomingEventsQueryOpts = (
	publicOnly: boolean = true
) => queryOptions({
	queryKey: [
		"events",
		"upcoming",
		...(publicOnly) ? ["public-only"] : [],
	],
	queryFn: () => listUpcomingEvents(publicOnly),
})

export const previousEventsQueryOpts =
	queryOptions({
		queryKey: ["events", "previous"],
		queryFn: listPreviousEvents,
	})

export const createEventByIdOpts = (eventId: ID) =>
	queryOptions({
		queryKey: ["events", eventId],
		queryFn: () => eventById(eventId)
	})

export const createVenueByIdQueryOptions = (venueId: ID) =>
	queryOptions({
		queryKey: ["venues", venueId],
		queryFn: () => venueById(venueId)
	})
