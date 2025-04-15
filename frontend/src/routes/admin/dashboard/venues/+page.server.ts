import { userRoles, hasSomeRole } from "$lib/auth";
import { listVenues } from "$lib/venue";
import { redirect } from "@sveltejs/kit";
import { userSession } from "$lib/auth";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const user = await userSession(fetch)
  const roles = await userRoles(fetch, user.id)

  if (!hasSomeRole(roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records: venues } = await listVenues(fetch)

  return { venues }
}
