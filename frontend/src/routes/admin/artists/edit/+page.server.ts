import { error, redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { artistById } from "$lib/artist";
import { listGenres } from "$lib/artist";
import { APIError } from "$lib/api";
import { hasPermissions } from "$lib/auth";

export const load: PageServerLoad = async ({ locals, url }) => {
  if (!hasPermissions(locals.permissions, ["view:artist", "edit:artist"])) {
    return redirect(302, "/auth/login")
  }

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
