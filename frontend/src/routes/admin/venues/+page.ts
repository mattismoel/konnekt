import { listVenues } from "$lib/features/venue/venue";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
  const { records: venues } = await listVenues(fetch)

  return { venues }
}
