import { getUpcomingEvents } from "$lib/features/event.remote";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
	const { items: upcomingEvents } = await getUpcomingEvents(undefined);

	return {
		upcomingEvents: upcomingEvents
	};
};
