import type { Event } from "../events/event";

/**
 * @description Extracts all genres of an event's concerts, and sorts them alphabetically.
 */
export const eventGenres = (event: Event): string[] => {
	const genres = new Set<string>();
	event.concerts.forEach(({ artist }) => {
		artist.genres.forEach((genre) => genres.add(genre))
	})

	return [...genres].sort((a, b) => a.localeCompare(b))
}
