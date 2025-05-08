import { venueById } from "$lib/features/venue/venue";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ params, fetch }) => {
	const venue = await venueById(fetch, parseInt(params.venueId))

	return { venue }
}
