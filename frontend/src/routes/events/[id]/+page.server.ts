import type { PageServerLoad } from "./$types";
import { eventById, listEvents } from "$lib/event";
import { startOfToday } from "date-fns";

export const load: PageServerLoad = async ({ params }) => {
  const id = parseInt(params.id)

  const event = await eventById(id)

  const recommendedEventsResult = await listEvents(new URLSearchParams({
    from: startOfToday().toISOString()
  }))

  return {
    event,
    recommendedEvents: recommendedEventsResult.records,
  }
}
