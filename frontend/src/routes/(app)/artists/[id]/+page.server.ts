import type { PageServerLoad } from "./$types";
import { artistById } from "$lib/artist";
import { artistEvents } from "$lib/event";

export const load: PageServerLoad = async ({ params, fetch }) => {
  const artist = await artistById(fetch, parseInt(params.id))

  const { records: events } = await artistEvents(fetch, parseInt(params.id))

  return { artist, events }
}
