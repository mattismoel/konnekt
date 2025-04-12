import { listUpcomingEvents } from "$lib/event"
import type { PageServerLoad } from "./$types"

export const load: PageServerLoad = async ({ fetch }) => {
  const { records } = await listUpcomingEvents(fetch)

  return {
    events: records
  }
}

