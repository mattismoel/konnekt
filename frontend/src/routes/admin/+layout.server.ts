import { redirect } from "@sveltejs/kit";
import { userPermissions } from "$lib/features/auth/permission";
import { userRoles, hasAllRoles } from "$lib/features/auth/role";
import { userSession } from "$lib/features/auth/user";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ fetch }) => {
  const user = await userSession(fetch)
  const roles = await userRoles(fetch, user.id)
  const permissions = await userPermissions(fetch, user.id)

  if (!hasAllRoles(roles, ["member"])) {
    return redirect(302, "/auth/login")
  }

  return { user, roles, permissions }
}
