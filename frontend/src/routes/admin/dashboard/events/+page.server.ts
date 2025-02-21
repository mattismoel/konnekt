import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { startOfToday } from "date-fns";
import type { PageServerLoad } from "./$types";
import { error } from "@sveltejs/kit";
import { eventSchema } from "$lib/event";
import { createListResult } from "$lib/list-result";

const UPCOMING_EVENTS_LIMIT: number = 5

export const load: PageServerLoad = async () => {
  let res = await fetch(`${PUBLIC_BACKEND_URL}/events?` + new URLSearchParams({
    filter: `from=${startOfToday().toISOString()}`,
    limit: UPCOMING_EVENTS_LIMIT.toString(),
    order: `date,asc`
  }))

  if (!res.ok) {
    return error(500, "Could not get events")
  }

  const upcomingEventsResult = createListResult(eventSchema).parse(await res.json())

  return {
    upcomingEvents: upcomingEventsResult.records,
  }
}
