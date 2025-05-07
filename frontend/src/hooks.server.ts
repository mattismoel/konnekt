import { APIError, apiErrorSchema } from "$lib/api";
import { redirect, type HandleFetch } from "@sveltejs/kit";

export const handleFetch: HandleFetch = async ({ request, fetch }) => {
	const response = await fetch(request, {
		headers: {
			...request.headers,
			"Content-Type": "applicaion/json",
		},
	})

	if (response.ok) {
		return response
	}

	if (response.status === 401) {
		return redirect(302, "/auth/login")
	}

	const err = apiErrorSchema.parse(await response.json())

	throw new APIError(response.status, err.message, "")
}
