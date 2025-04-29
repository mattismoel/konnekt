import { listGenres } from "$lib/features/artist/genre";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
	const { records: genres } = await listGenres(fetch)

	return { genres }
}
