import type { PageServerLoad } from "./$types";
import { startOfToday } from "date-fns";
import { eventById, listEvents } from "$lib/event";

export const load: PageServerLoad = async ({ params, fetch }) => {
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
