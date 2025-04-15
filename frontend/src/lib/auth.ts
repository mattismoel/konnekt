import { z } from "zod"

import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { requestAndParse } from "./api"
import { createUrl } from "./url";

export const SESSION_COOKIE_NAME = "konnekt-session"
export const SESSION_COOKIE_LIFETIME_MILLIS = 30 * 60000 * 60 * 24 // 30 days.

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

export const permissionTypes = z.union([
	z.literal("view:event"),
	z.literal("edit:event"),
	z.literal("delete:event"),

	z.literal("view:artist"),
	z.literal("edit:artist"),
	z.literal("delete:artist"),

	z.literal("view:venue"),
	z.literal("edit:venue"),
	z.literal("delete:venue"),

	z.literal("view:role"),
	z.literal("edit:role"),
	z.literal("delete:role"),

	z.literal("view:genre"),
	z.literal("edit:genre"),
	z.literal("delete:genre"),

	z.literal("view:permission")
])

export const roleTypes = z.union([
	z.literal("admin"),
	z.literal("member"),
	z.literal("event-management"),
	z.literal("booking"),
	z.literal("public-relations"),
	z.literal("visual-identity"),
])

export type RoleType = z.infer<typeof roleTypes>

export const roleSchema = z.object({
	id: z.number().positive(),
	name: roleTypes,
	displayName: z.string(),
	description: z.string()
})

export type Role = z.infer<typeof roleSchema>

export const permissionSchema = z.object({
	id: z.number().int().positive(),
	name: permissionTypes,
	displayName: z.string().nonempty(),
	description: z.string().nonempty(),
})

export type PermissionType = z.infer<typeof permissionTypes>
export type Permission = z.infer<typeof permissionSchema>

export const userRoles = async (fetchFn: typeof fetch, userId: number): Promise<Role[]> => {
	const roles = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/roles/${userId}`),
		roleSchema.array(),
		"Could not fetch user roles",
	)

	return roles
}

export const userPermissions = async (fetchFn: typeof fetch, userId: number): Promise<Permission[]> => {
	const permissions = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/permissions/${userId}`),
		permissionSchema.array(),
		"Could not fetch user permissions",
	)

	return permissions
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

export const hasPermissions = (userPermissions: Permission[], permNames: PermissionType[]): boolean => {
	return permNames.every(perm => userPermissions.some(userPerm => userPerm.name === perm))
}
