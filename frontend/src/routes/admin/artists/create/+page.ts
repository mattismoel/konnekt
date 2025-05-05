import { listGenres } from "$lib/features/artist/genre";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	const { records: genres } = await listGenres(fetch)

	return { genres }
}
