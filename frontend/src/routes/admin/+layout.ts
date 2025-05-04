import type { LayoutLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { hasAllTeams } from "$lib/features/auth/team";
import { memberSession } from "$lib/features/auth/member";
import { authStore } from "$lib/auth.svelte";


export const load: LayoutLoad = async ({ fetch }) => {
  const member = await memberSession(fetch)

  if (!hasAllTeams(member.teams, ["member"])) {
    return redirect(302, "/auth/login")
  }


  authStore.auth = {
    member,
    permissions: member.permissions,
    teams: member.teams,
  }
  console.log(authStore.permissions)
}
