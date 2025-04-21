import type { PageServerLoad } from "./$types";
import { error, redirect } from "@sveltejs/kit";

import { artistById } from "$lib/features/artist/artist";
import { listGenres } from "$lib/features/artist/genre";
import { hasPermissions } from "$lib/features/auth/permission";
import { memberSession } from "$lib/features/auth/member";

export const load: PageServerLoad = async ({ fetch, url }) => {
  const member = await memberSession(fetch)

  if (!hasPermissions(member.permissions, ["view:artist", "edit:artist"])) {
    return redirect(302, "/auth/login")
  }

  const id = url.searchParams.get("id")
  if (!id) return error(500, "No artist ID")

  const { records: genres } = await listGenres(fetch)

  const artist = await artistById(fetch, parseInt(id))

  return { genres, artist }
}
