import type { PageServerLoad } from "./$types"
import { listUpcomingEvents } from "$lib/event"

export const load: PageServerLoad = async () => {
  const { records } = await listUpcomingEvents()
  console.log(records)
  return {
    events: records
  }
}

