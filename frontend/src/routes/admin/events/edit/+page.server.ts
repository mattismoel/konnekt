import type { PageServerLoad } from "./$types";

import { listArtists } from "$lib/features/artist/artist";
import { eventById } from "$lib/features/event/event";
import { listVenues } from "$lib/features/venue/venue";
import { error } from "@sveltejs/kit";

export const load: PageServerLoad = async ({ url, fetch }) => {
	const { records: venues } = await listVenues(fetch)
	const { records: artists } = await listArtists(fetch)

	const eventId = url.searchParams.get("id")
	if (!eventId) error(500, "No event ID found")

	const event = await eventById(fetch, parseInt(eventId))

	return {
		venues,
		artists,
		event,
	}
}
