import type { LayoutServerLoad } from "./$types";

import { hasAllTeams, memberTeams } from "$lib/features/auth/team";
import { memberSession } from "$lib/features/auth/member";
import { authStore } from "$lib/auth.svelte";
import { redirect } from "@sveltejs/kit";
import { memberPermissions } from "$lib/features/auth/permission";

export const load: LayoutServerLoad = async ({ fetch }) => {
  try {
    const member = await memberSession(fetch)
    const teams = await memberTeams(fetch, member.id)
    const permissions = await memberPermissions(fetch, member.id)

    if (!hasAllTeams(teams, ["member"]) || !member.active) {
      return redirect(302, "/auth/login")
    }

    authStore.auth = { member, permissions, teams }
  } catch (e) {
    authStore.auth = { member: null, permissions: [], teams: [] }
    return redirect(302, "/auth/login")
  }
}
