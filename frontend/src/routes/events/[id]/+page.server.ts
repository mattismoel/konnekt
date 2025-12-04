import { getEvent } from "$lib/features/event.remote";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params: { id } }) => {
	const event = await getEvent(id);
	return { event };
};
