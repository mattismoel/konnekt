import type { Handle } from "@sveltejs/kit";

import { PUBLIC_BACKEND_URL } from "$env/static/public";

import { listUserPermissions, listUserRoles } from "$lib/auth";
import { userSchema } from "$lib/user";
import { tryCatch } from "$lib/error";

export const handle: Handle = async ({ event, resolve }) => {
	let res = await fetch(`${PUBLIC_BACKEND_URL}/auth/session`, {
		credentials: "include",
		headers: event.request.headers
	})

	if (!res.ok) {
		event.locals.user = null
	} else {
		event.locals.user = userSchema.parse(await res.json())
	}

	if (!event.locals.user) return await resolve(event)

	const rolesResp = await tryCatch(listUserRoles(event.locals.user?.id, { headers: event.request.headers }))
	if (rolesResp.error) {
		event.locals.roles = []
		return await resolve(event)
	}

	const permsResp = await tryCatch(listUserPermissions(event.locals.user.id, { headers: event.request.headers }))
	if (permsResp.error) {
		event.locals.roles = []
		return await resolve(event)
	}

	event.locals.roles = rolesResp.data
	event.locals.permissions = permsResp.data

	return await resolve(event)
}
