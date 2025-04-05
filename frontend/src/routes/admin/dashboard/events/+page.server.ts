import type { PageServerLoad } from "./$types";
import { startOfToday } from "date-fns";
import { listEvents } from "$lib/event";
import { hasSomeRole } from "$lib/auth";
import { redirect } from "@sveltejs/kit";


export const load: PageServerLoad = async ({ locals, request }) => {
  if (!hasSomeRole(locals.roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const upcomingEventsResult = await listEvents(new URLSearchParams({
    filter: `from_date>=${startOfToday().toISOString()}`
  }))

  const previousEventsResult = await listEvents(new URLSearchParams({
    filter: `from_date<${startOfToday().toISOString()}`
  }))

  return {
    upcomingEvents: upcomingEventsResult.records,
    previousEvents: previousEventsResult.records,
  }
}
