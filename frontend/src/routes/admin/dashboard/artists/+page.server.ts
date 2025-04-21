import { listArtists } from "$lib/features/artist/artist";
import { hasSomeRole } from "$lib/features/auth/role";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { memberSession } from "$lib/features/auth/member";

export const load: PageServerLoad = async ({ fetch }) => {
  const member = await memberSession(fetch)

  if (!hasSomeRole(member.roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records } = await listArtists(fetch)

  return {
    artists: records
  }
}
