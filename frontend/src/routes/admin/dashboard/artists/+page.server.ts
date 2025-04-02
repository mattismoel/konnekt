import { listArtists } from "$lib/artist";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async () => {
  const { records } = await listArtists()

  return {
    artists: records
  }
}
