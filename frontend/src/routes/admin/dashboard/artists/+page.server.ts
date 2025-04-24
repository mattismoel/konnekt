import { listArtists } from "$lib/features/artist/artist";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { memberSession } from "$lib/features/auth/member";
import { hasSomeTeam } from "$lib/features/auth/team";

export const load: PageServerLoad = async ({ fetch }) => {
  const member = await memberSession(fetch)

  if (!hasSomeTeam(member.teams, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records } = await listArtists(fetch)

  return {
    artists: records
  }
}
