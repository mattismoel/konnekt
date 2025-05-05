import { venueById } from "$lib/features/venue/venue";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params, fetch }) => {
	const venue = await venueById(fetch, parseInt(params.venueId))

	return { venue }
}
