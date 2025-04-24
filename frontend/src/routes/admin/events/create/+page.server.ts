import { listArtists } from "$lib/features/artist/artist";
import { listVenues } from "$lib/features/venue/venue";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
	const { records: venues } = await listVenues(fetch)
	const { records: artists } = await listArtists(fetch)

	return { venues, artists }
}
