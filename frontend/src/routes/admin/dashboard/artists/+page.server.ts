import { listArtists } from "$lib/features/artist/artist";
import { hasSomeRole } from "$lib/features/auth/role";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { userSession } from "$lib/features/auth/user";

export const load: PageServerLoad = async ({ fetch }) => {
  const user = await userSession(fetch)

  if (!hasSomeRole(user.roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records } = await listArtists(fetch)

  return {
    artists: records
  }
}
