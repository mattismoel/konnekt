import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { artistById } from "$lib/artist";
import { listGenres } from "$lib/genre";
import { APIError } from "$lib/error";

export const load: PageServerLoad = async ({ url }) => {
  try {
    const id = url.searchParams.get("id")

    const genresResult = await listGenres()

    if (!id) return { artist: null, genres: genresResult.records }

    const artist = await artistById(parseInt(id))

    return {
      genres: genresResult.records,
      artist
    }
  } catch (e) {
    if (e instanceof APIError) return error(e.status, e.message)
    throw e
  }
}
