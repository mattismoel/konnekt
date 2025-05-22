import { queryOptions } from "@tanstack/react-query";
import { artistById, listArtists } from "./artist";
import { artistEvents, listUpcomingEvents } from "../event/event";
import { listGenres } from "./genre";
import { eventsArtists } from "./events-artists";
import type { ID } from "@/lib/api";

export const artistsQueryOpts =
	queryOptions({
		queryKey: ["artists"],
		queryFn: () => listArtists()
	})

export const createArtistByIdOpts = (artistId: ID) =>
	queryOptions({
		queryKey: ["artists", artistId],
		queryFn: () => artistById(artistId)
	})

export const createArtistEventsOpts = (artistId: ID) =>
	queryOptions({
		queryKey: ["events", "artist-id", artistId],
		queryFn: () => artistEvents(artistId)
	})

export const genresQueryOpts =
	queryOptions({
		queryKey: ["genres"],
		queryFn: () => listGenres()
	})

export const upcomingArtistsQueryOpts =
	queryOptions({
		queryKey: ["artists", "upcoming"],
		queryFn: async () => {
			const { records: upcomingEvents } = await listUpcomingEvents()

			const upcomingArtists = eventsArtists(upcomingEvents)
			return upcomingArtists
		}
	})
