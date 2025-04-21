import { memberSession } from "$lib/features/auth/member";
import { hasSomeRole } from "$lib/features/auth/role";
import { listVenues } from "$lib/features/venue/venue";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
  const member = await memberSession(fetch)

  if (!hasSomeRole(member.roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records: venues } = await listVenues(fetch)

  return { venues }
}
