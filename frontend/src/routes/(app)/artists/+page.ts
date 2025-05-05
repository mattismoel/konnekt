import type { PageLoad } from "./$types";
import { listUpcomingEvents } from "$lib/features/event/event";
import { eventsArtists } from "$lib/features/artist/artist";

export const load: PageLoad = async ({ fetch }) => {
  const { records: upcomingEvents } = await listUpcomingEvents(fetch)

  let artists = eventsArtists(upcomingEvents)

  return {
    artists,
  }
}
