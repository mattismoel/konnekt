import type { PageServerLoad } from "./$types";
import { artistById } from "$lib/features/artist/artist";
import { artistEvents } from "$lib/features/event/event";

export const load: PageServerLoad = async ({ params, fetch }) => {
  const artist = await artistById(fetch, parseInt(params.id))

  const { records: events } = await artistEvents(fetch, parseInt(params.id))

  return { artist, events }
}
