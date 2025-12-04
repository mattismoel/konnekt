import { getUpcomingEvents } from "$lib/features/event.remote";
import { getLandingImages } from "$lib/features/landing.remote";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
	const { items: landingImages } = await getLandingImages(undefined);
	const { items: upcomingEvents } = await getUpcomingEvents(undefined);

	return {
		landingImages,
		upcomingEvents: upcomingEvents
	};
};
