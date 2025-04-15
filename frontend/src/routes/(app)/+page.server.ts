import { startOfToday } from "date-fns";

import type { PageServerLoad } from "./$types";

import { listUpcomingEvents } from "$lib/features/event/event";

export const load: PageServerLoad = async ({ fetch }) => {
	const { records } = await listUpcomingEvents(fetch)
	return { events: records }
}
