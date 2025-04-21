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

export const userRoles = async (fetchFn: typeof fetch, userId: number): Promise<Role[]> => {
	const roles = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/roles/${userId}`),
		roleSchema.array(),
		"Could not fetch user roles",
	)

	return roles
}


/**
 * @description Checks whether or not a given user has all required input roles.
 */
export const hasAllRoles = (userRoles: Role[], roleNames: RoleType[]): boolean => {
	return roleNames.every(role => userRoles.some(userRole => userRole.name === role))
}

export const hasSomeRole = (userRoles: Role[], roleNames: RoleType[]): boolean => {
	return roleNames.some(role => userRoles.some(userRole => userRole.name === role))
}
