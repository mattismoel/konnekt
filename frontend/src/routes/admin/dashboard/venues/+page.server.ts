import { hasSomeRole } from "$lib/auth";
import { listVenues } from "$lib/venue";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals, request }) => {
  if (!hasSomeRole(locals.roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records } = await listVenues({
    credentials: "include",
    headers: request.headers,
  })

  return {
    venues: records
  }
}
