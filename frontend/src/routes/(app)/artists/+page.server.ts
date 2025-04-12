import type { PageServerLoad } from "./$types";
import { listUpcomingEvents } from "$lib/event";
import { listArtists } from "$lib/artist";
import { removeDuplicates } from "$lib/array";

export const load: PageServerLoad = async ({ fetch }) => {
  let currentArtists = await listUpcomingEvents(fetch)
    .then(({ records }) => records
      .flatMap(({ concerts }) => concerts)
      .map(({ artist }) => artist)
    )

  currentArtists = removeDuplicates(currentArtists)

  const { records: artists } = await listArtists(fetch)

  return {
    currentArtists,
    artists,
  }
}
