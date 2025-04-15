import { listUpcomingEvents } from "$lib/features/event/event"
import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async ({ fetch }) => {
  const { records } = await listUpcomingEvents(fetch)

  return {
    events: records
  }
}

