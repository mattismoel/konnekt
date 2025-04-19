import { listArtists } from "$lib/features/artist/artist";
import { hasSomeRole, userRoles } from "$lib/features/auth/role";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { userSession } from "$lib/features/auth/user";

export const load: PageServerLoad = async ({ fetch }) => {
  const { id: userId } = await userSession(fetch)

  const roles = await userRoles(fetch, userId)


  if (!hasSomeRole(roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records } = await listArtists(fetch)

  return {
    artists: records
  }
}
