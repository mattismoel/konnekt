import type { LayoutLoad } from "./$types";

import { hasAllTeams } from "$lib/features/auth/team";
import { memberSession } from "$lib/features/auth/member";
import { authStore } from "$lib/auth.svelte";
import { redirect } from "@sveltejs/kit";

export const load: LayoutLoad = async ({ fetch }) => {
  try {
    const member = await memberSession(fetch)
    if (!hasAllTeams(member.teams, ["member"])) {
      return redirect(302, "/auth/login")
    }

    authStore.auth = {
      member,
      permissions: member.permissions,
      teams: member.teams,
    }
  } catch (e) {
    return redirect(302, "/auth/login")
  }
}
