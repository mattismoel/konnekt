import type { PageServerLoad } from "./$types"
import { listEvents } from "$lib/event"
import { startOfToday } from "date-fns"

export const load: PageServerLoad = async () => {
  const result = await listEvents(new URLSearchParams({
    from: startOfToday().toISOString(),
  }))

  return {
    events: result.records
  }
}

