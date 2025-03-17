import { redirect } from "@sveltejs/kit";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import type { LayoutServerLoad } from "./$types";
import { roleSchema } from "$lib/auth";

export const load: LayoutServerLoad = async ({ locals, request }) => {
  let res = await fetch(`${PUBLIC_BACKEND_URL}/users/roles/${locals.user?.id}`, {
    credentials: "include",
    headers: request.headers
  })

  if (!res.ok) {
    return redirect(307, "/auth/login")
  }

  const roles = roleSchema.array().parse(await res.json())

  if (!roles.some(role => role.name === "admin")) {
    return redirect(307, "/auth/login")
  }
}
