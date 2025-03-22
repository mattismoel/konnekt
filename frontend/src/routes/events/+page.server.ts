import type { PageServerLoad } from "./$types"
import { startOfToday } from "date-fns"
import { listEvents } from "$lib/event"

export const load: PageServerLoad = async () => {
  const result = await listEvents(new URLSearchParams({
    "filter": [
      "from_date" + ">=" + startOfToday().toISOString(),
    ].join(","),
    "order_by": "from_date desc"
  }))

  return {
    events: result.records
  }
}

