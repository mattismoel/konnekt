import { listArtists } from "$lib/artist";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ request }) => {
  const artistsResult = await listArtists(new URLSearchParams())

  return {
    artists: artistsResult.records
  }
}
