import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { userSchema } from "$lib/user";
import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
	console.log(`${PUBLIC_BACKEND_URL}`)
	const res = await fetch(`${PUBLIC_BACKEND_URL}/auth/session`, {
		credentials: "include",
		headers: event.request.headers
	})

	if (!res.ok) {
		event.locals.user = null
	} else {
		event.locals.user = userSchema.parse(await res.json())
	}

	const response = await resolve(event)

	return response
}
