import type { PageServerLoad } from "./$types";
import { listUpcomingEvents } from "$lib/event";
import { listArtists } from "$lib/artist";
import { removeDuplicates } from "$lib/array";

export const load: PageServerLoad = async ({ request }) => {
  let currentArtists = await listUpcomingEvents()
    .then(({ records }) => records
      .flatMap(({ concerts }) => concerts)
      .map(({ artist }) => artist)
    )

  currentArtists = removeDuplicates(currentArtists)

  console.log(currentArtists.map(a => ({ id: a.id, name: a.name })))

  const { records: artists } = await listArtists(new URLSearchParams())

  return {
    currentArtists,
    artists,
  }
}
