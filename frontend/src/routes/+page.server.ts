import { listEvents } from "$lib/event";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ }) => {
	const { records } = await listEvents(new URLSearchParams())

	return { events: records }
}
