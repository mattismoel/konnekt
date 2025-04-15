import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { requestAndParse } from "$lib/api"
import { createUrl } from "$lib/url"
import { z } from "zod"

export const userSchema = z.object({
	id: z.number().positive(),
	email: z.string().email(),
	firstName: z.string(),
	lastName: z.string()
})

export type User = z.infer<typeof userSchema>

export const userSession = async (fetchFn: typeof fetch) => {
	const user = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/session`),
		userSchema,
		"Could not fetch user session",
	)

	return user
}
