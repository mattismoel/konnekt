import type { PageServerLoad } from "./$types"
import { startOfToday } from "date-fns"
import { listEvents } from "$lib/event"

export const load: PageServerLoad = async () => {
  const result = await listEvents(new URLSearchParams({
    from: startOfToday().toISOString(),
  }))

  return {
    events: result.records
  }
}

