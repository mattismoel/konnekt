import type { PageServerLoad } from "./$types";

import { listEvents } from "$lib/event";

export const load: PageServerLoad = async ({ }) => {
	const { records } = await listEvents(new URLSearchParams())

	return { events: records }
}
