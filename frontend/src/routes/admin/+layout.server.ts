import { redirect } from "@sveltejs/kit";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import type { LayoutServerLoad } from "./$types";
import { hasAllRoles, roleSchema } from "$lib/auth";

export const load: LayoutServerLoad = async ({ locals, request }) => {
  if (!locals.user) return redirect(302, "/auth/login")

  if (!hasAllRoles(locals.roles, ["member"])) {
    return redirect(302, "/auth/login")
  }

  return {
    user: locals.user,
    roles: locals.roles,
    permissions: locals.permissions,
  }
}
