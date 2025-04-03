import type { Handle } from "@sveltejs/kit";

import { PUBLIC_BACKEND_URL } from "$env/static/public";

import { roleSchema } from "$lib/auth";
import { userSchema } from "$lib/user";

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

	res = await fetch(`${PUBLIC_BACKEND_URL}/auth/roles/${event.locals.user?.id}`, {
		credentials: "include",
		headers: event.request.headers
	})

	if (!res.ok) {
		event.locals.roles = []
	} else {
		event.locals.roles = roleSchema.array().parse(await res.json())
	}

	console.log(event.locals.user, event.locals.roles)

	const response = await resolve(event)

	return response
}
