import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { hasAllTeams } from "$lib/features/auth/team";
import { memberSession } from "$lib/features/auth/member";


export const load: LayoutServerLoad = async ({ fetch }) => {
  const member = await memberSession(fetch)

  if (!hasAllTeams(member.teams, ["member"])) {
    return redirect(302, "/auth/login")
  }

  return { member }
}
