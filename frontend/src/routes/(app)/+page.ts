import type { PageLoad } from "./$types";

import { listUpcomingEvents } from "$lib/features/event/event";
import { eventsArtists } from "$lib/features/artist/artist";

export const load: PageLoad = async ({ fetch }) => {
	const { records: events } = await listUpcomingEvents(fetch)

	const artists = eventsArtists(events)

	return { events, artists }
}
