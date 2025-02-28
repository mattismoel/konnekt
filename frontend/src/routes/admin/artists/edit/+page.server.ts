import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { createListResult } from "$lib/list-result";
import { artistSchema } from "$lib/artist";
import { genreSchema } from "$lib/genre";

export const load: PageServerLoad = async ({ url }) => {
  const id = url.searchParams.get("id")

  let res = await fetch(`${PUBLIC_BACKEND_URL}/genres`)
  if (!res.ok) {
    return error(500, "Could not load genres")
  }

  const genres = createListResult(genreSchema).parse(await res.json()).records

  if (!id) return { artist: null, genres }

  res = await fetch(`${PUBLIC_BACKEND_URL}/artists/${id}`)
  if (!res.ok) {
    return error(500, "Could not load artist")
  }

  const artist = artistSchema.parse(await res.json())

  return {
    genres,
    artist
  }
}
