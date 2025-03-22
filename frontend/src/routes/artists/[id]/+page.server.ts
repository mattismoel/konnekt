import type { PageServerLoad } from "./$types";
import { artistById } from "$lib/artist";
import { listEvents } from "$lib/event";
import { startOfToday } from "date-fns";

export const load: PageServerLoad = async ({ params }) => {
  const artist = await artistById(parseInt(params.id))
  const artistEventsResult = await listEvents(new URLSearchParams({
    "filter": [
      "from_date>=" + startOfToday().toISOString(),
      "artist_id=" + parseInt(params.id)
    ].join(",")
  }))

  return {
    artist,
    events: artistEventsResult.records,
  }
}
