import { listArtists } from "$lib/features/artist/artist";
import { hasSomeRole } from "$lib/features/auth/role";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ locals }) => {
  if (!hasSomeRole(locals.roles, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records } = await listArtists()

  return {
    artists: records
  }
}
