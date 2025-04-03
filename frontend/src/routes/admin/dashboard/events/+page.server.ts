import type { PageServerLoad } from "./$types";
import { startOfToday } from "date-fns";
import { listEvents } from "$lib/event";
import { hasSomeRole } from "$lib/auth";
import { redirect } from "@sveltejs/kit";

const UPCOMING_EVENTS_LIMIT: number = 5

export const load: PageServerLoad = async ({ locals, request }) => {
  if (!hasSomeRole(locals.roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const upcomingEventsResult = await listEvents(new URLSearchParams({
    filter: `from_date>=${startOfToday().toISOString()}`,
    limit: UPCOMING_EVENTS_LIMIT.toString(),
  }))

  return {
    upcomingEvents: upcomingEventsResult.records,
  }
}
