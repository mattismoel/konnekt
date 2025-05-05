import type { PageLoad } from "./$types";
import { startOfToday } from "date-fns";
import { eventById, listEvents } from "$lib/features/event/event";

export const load: PageLoad = async ({ params, fetch }) => {
  const id = parseInt(params.id)

  const event = await eventById(fetch, id)

  const { records: recommendedEvents } = await listEvents(fetch, {
    filter: [
      "id" + "!=" + id,
      "from_date" + ">=" + startOfToday().toISOString(),
    ]
  })

  return { event, recommendedEvents }
}
