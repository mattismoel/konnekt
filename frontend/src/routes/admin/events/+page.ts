import { listPreviousEvents, listUpcomingEvents } from "$lib/features/event/event";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
  const { records: upcomingEvents } = await listUpcomingEvents(fetch)
  const { records: previousEvents } = await listPreviousEvents(fetch)

  return { upcomingEvents, previousEvents }
}
