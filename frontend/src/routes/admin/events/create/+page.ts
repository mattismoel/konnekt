import { listArtists } from "$lib/features/artist/artist";
import { listVenues } from "$lib/features/venue/venue";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	const { records: venues } = await listVenues(fetch)
	const { records: artists } = await listArtists(fetch)

	return { venues, artists }
}
