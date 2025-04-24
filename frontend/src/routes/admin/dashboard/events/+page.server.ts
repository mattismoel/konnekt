import { listPreviousEvents, listUpcomingEvents } from "$lib/features/event/event";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const { records: upcomingEvents } = await listUpcomingEvents(fetch)
  const { records: previousEvents } = await listPreviousEvents(fetch)

  return { upcomingEvents, previousEvents }
}
