import type { PageLoad } from "./$types";

import { listArtists } from "$lib/features/artist/artist";
import { eventById } from "$lib/features/event/event";
import { listVenues } from "$lib/features/venue/venue";

export const load: PageLoad = async ({ params, fetch }) => {
	const { records: venues } = await listVenues(fetch)
	const { records: artists } = await listArtists(fetch)

	const event = await eventById(fetch, parseInt(params.eventId))

	return {
		venues,
		artists,
		event,
	}
}
