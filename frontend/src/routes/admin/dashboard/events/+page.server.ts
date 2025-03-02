import { startOfToday } from "date-fns";
import type { PageServerLoad } from "./$types";
import { listEvents } from "$lib/event";

const UPCOMING_EVENTS_LIMIT: number = 5

export const load: PageServerLoad = async () => {
  const upcomingEventsResult = await listEvents(new URLSearchParams({
    filter: `from=${startOfToday().toISOString()}`,
    limit: UPCOMING_EVENTS_LIMIT.toString(),
    order: `date,asc`
  }))

  return {
    upcomingEvents: upcomingEventsResult.records,
  }
}
