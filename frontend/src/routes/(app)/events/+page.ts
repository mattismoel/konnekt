import { listUpcomingEvents } from "$lib/features/event/event"
import type { PageLoad } from "./$types"

export const load: PageLoad = async ({ fetch }) => {
  const { records } = await listUpcomingEvents(fetch)

  return {
    events: records
  }
}

