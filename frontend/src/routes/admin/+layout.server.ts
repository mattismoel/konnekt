import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { hasAllRoles } from "$lib/features/auth/role";
import { memberSession } from "$lib/features/auth/member";


export const load: LayoutServerLoad = async ({ fetch }) => {
  const member = await memberSession(fetch)

  if (!hasAllRoles(member.roles, ["member"])) {
    return redirect(302, "/auth/login")
  }

  return { member }
}
