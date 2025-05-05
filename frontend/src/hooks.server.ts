import { cleanUrl } from "$lib/url";
import { redirect, type Handle, type HandleFetch } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
	if (cleanUrl(event.url.pathname) === "/admin") {
		redirect(302, "/admin/events")
	}

	return await resolve(event)
}

export const handleFetch: HandleFetch = async ({ request, fetch }) => {
	return fetch(request, {
		headers: {
			"Content-Type": "application/json",
			...request.headers
		}
	})
}
