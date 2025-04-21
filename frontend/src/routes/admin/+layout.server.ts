import type { LayoutServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";

import { hasAllRoles } from "$lib/features/auth/role";
import { userSession } from "$lib/features/auth/user";


export const load: LayoutServerLoad = async ({ fetch }) => {
  const user = await userSession(fetch)

  if (!hasAllRoles(user.roles, ["member"])) {
    return redirect(302, "/auth/login")
  }

  return { user }
}
