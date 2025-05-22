import { queryOptions } from "@tanstack/react-query";
import { listVenues, venueById } from "./venue";
import { eventById, listPreviousEvents, listUpcomingEvents } from "./event";
import type { ID } from "@/lib/api";

export const venuesQueryOpts =
	queryOptions({
		queryKey: ["venues"],
		queryFn: () => listVenues()
	})

export const upcomingEventsQueryOpts =
	queryOptions({
		queryKey: ["events", "upcoming"],
		queryFn: listUpcomingEvents,
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
