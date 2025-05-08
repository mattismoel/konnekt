import type { PageLoad } from "./$types";

import { artistById } from "$lib/features/artist/artist";
import { listGenres } from "$lib/features/artist/genre";

export const load: PageLoad = async ({ params, fetch }) => {
  const { records: genres } = await listGenres(fetch)

  const artist = await artistById(fetch, parseInt(params.artistId))

  return { genres, artist }
}
