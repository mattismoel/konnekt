import type { PageServerLoad } from "./$types";
import { listArtists } from "$lib/artist";

export const load: PageServerLoad = async ({ request }) => {
  const artistsResult = await listArtists(new URLSearchParams())

  return {
    artists: artistsResult.records
  }
}
