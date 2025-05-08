import { eventsArtists, listArtists } from "$lib/features/artist/artist";
import type { PageLoad } from "./$types";
import { listUpcomingEvents } from "$lib/features/event/event";

export const load: PageLoad = async ({ fetch }) => {
  const { records: upcomingEvents } = await listUpcomingEvents(fetch)

  const upcomingArtists = eventsArtists(upcomingEvents)

  const { records: artists } = await listArtists(fetch)

  return {
    artists,
    upcomingArtists,
  }
}
