import { useQuery } from "@tanstack/react-query"
import { artistEvents, eventById, listPreviousEvents, listUpcomingEvents } from "./event"
import { eventsArtists } from "./events-artists"
import { artistById, listArtists } from "./artist"
import { listVenues } from "./venue"
import { listGenres } from "./genre"

export const useListUpcomingArtists = () => {
	return useQuery({
		queryKey: ["artists", "upcoming"],
		queryFn: async () => {
			const { records: upcomingEvents } = await listUpcomingEvents()

			const upcomingArtists = eventsArtists(upcomingEvents)
			return upcomingArtists
		}
	})
}

export const useArtists = () => {
	return useQuery({
		queryKey: ["artists"],
		queryFn: () => listArtists()
	})
}

export const useVenues = () => {
	return useQuery({
		queryKey: ["venues"],
		queryFn: () => listVenues()
	})
}

export const useListUpcomingEvents = () => {
	return useQuery({
		queryKey: ["events", "upcoming"],
		queryFn: listUpcomingEvents,
	})
}

export const useListPreviousEvents = () => {
	return useQuery({
		queryKey: ["events", "previous"],
		queryFn: listPreviousEvents,
	})
}

export const useArtistById = (id: number) => {
	return useQuery({
		queryKey: ["artists", id],
		queryFn: () => artistById(id)
	})
}

export const useArtistEvents = (id: number) => {
	return useQuery({
		queryKey: ["events", "artist-id", id],
		queryFn: () => artistEvents(id)
	})
}

export const useEventById = (eventId: number) => {
	return useQuery({
		queryKey: ["events", eventId],
		queryFn: () => eventById(eventId)
	})
}

export const useGenres = () => {
	return useQuery({
		queryKey: ["genres"],
		queryFn: () => listGenres()
	})
}
