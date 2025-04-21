import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { requestAndParse } from "$lib/api"
import { createListResult } from "$lib/query"
import { createUrl } from "$lib/url"
import { z } from "zod"
import { roleSchema } from "./role"
import { permissionSchema } from "./permission"

export const memberSchema = z.object({
	id: z.number().positive(),
	email: z.string().email(),
	firstName: z.string(),
	lastName: z.string(),

	roles: roleSchema.array(),
	permissions: permissionSchema.array(),
})

export type Member = z.infer<typeof memberSchema>

export const memberSession = async (fetchFn: typeof fetch) => {
	const member = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/session`),
		memberSchema,
		"Could not fetch member session",
	)

	return member
}

export const listMembers = async (fetchFn: typeof fetch) => {
	const result = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/members`),
		createListResult(memberSchema),
		"Could not fetch members",
	)

	return result
}
