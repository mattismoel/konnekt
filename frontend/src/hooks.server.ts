import type { HandleFetch } from "@sveltejs/kit";

export const handleFetch: HandleFetch = async ({ event, request, fetch }) => {
	return fetch(request, {
		headers: {
			"Content-Type": "application/json",
			...request.headers
		}
	})
}
