import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { requestAndParse } from "$lib/api"
import { createUrl } from "$lib/url"
import { z } from "zod"

export const roleTypes = z.union([
	z.literal("admin"),
	z.literal("member"),
	z.literal("event-management"),
	z.literal("booking"),
	z.literal("public-relations"),
	z.literal("visual-identity"),
	z.literal("team-management")
])

export type RoleType = z.infer<typeof roleTypes>

export const roleSchema = z.object({
	id: z.number().positive(),
	name: roleTypes,
	displayName: z.string(),
	description: z.string()
})

export type Role = z.infer<typeof roleSchema>

export const memberRoles = async (fetchFn: typeof fetch, memberId: number): Promise<Role[]> => {
	const roles = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/roles/${memberId}`),
		roleSchema.array(),
		"Could not fetch member roles",
	)

	return roles
}


/**
 * @description Checks whether or not a given member has all required input roles.
 */
export const hasAllRoles = (roles: Role[], reqRoleNames: RoleType[]): boolean => {
	return reqRoleNames.every(role => roles.some(r => r.name === role))
}

export const hasSomeRole = (roles: Role[], reqRoleNames: RoleType[]): boolean => {
	return reqRoleNames.some(role => roles.some(r => r.name === role))
}
