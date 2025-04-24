import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { artistById } from "$lib/features/artist/artist";
import { listGenres } from "$lib/features/artist/genre";
import { hasPermissions } from "$lib/features/auth/permission";
import { memberSession } from "$lib/features/auth/member";

export const load: PageServerLoad = async ({ params, fetch }) => {
  const member = await memberSession(fetch)

  if (!hasPermissions(member.permissions, ["view:artist", "edit:artist"])) {
    return redirect(302, "/auth/login")
  }

  const { records: genres } = await listGenres(fetch)

  const artist = await artistById(fetch, parseInt(params.artistId))

  return { genres, artist }
}
