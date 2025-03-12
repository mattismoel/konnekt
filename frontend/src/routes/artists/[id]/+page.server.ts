import { artistById } from "$lib/artist";
import { artistEvents } from "$lib/event";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
  const artist = await artistById(parseInt(params.id))
  const eventsResult = await artistEvents(parseInt(params.id))

  return {
    artist,
    events: eventsResult.records,
  }
}
