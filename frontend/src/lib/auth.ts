import { z } from "zod"

import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { APIError, apiErrorSchema } from "./error"

export const SESSION_COOKIE_NAME = "konnekt-session"
export const SESSION_COOKIE_LIFETIME_MILLIS = 30 * 60000 * 60 * 24 // 30 days.

export const roleSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	displayName: z.string(),
	description: z.string()
})

export type Role = z.infer<typeof roleSchema>

export const permissionSchema = z.object({
	id: z.number().int().positive(),
	name: z.string().nonempty(),
	displayName: z.string().nonempty(),
	description: z.string().nonempty(),
})

export type Permission = z.infer<typeof permissionSchema>
export const listUserRoles = async (userId: number, init?: RequestInit): Promise<Role[]> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/auth/roles/${userId}`, {
		credentials: "include",
		...init,
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not list user roles", err.message)
	}

	const roles = roleSchema.array().parse(await res.json())

	return roles
}

export const listUserPermissions = async (userId: number, init?: RequestInit): Promise<Permission[]> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/auth/permissions/${userId}`, {
		credentials: "include",
		...init,
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not list user roles", err.message)
	}

	const permissions = permissionSchema.array().parse(await res.json())

	return permissions
}

/**
 * @description Checks whether or not a given user has all required input roles.
 */
export const hasAllRoles = (userRoles: Role[], roleNames: string[]): boolean => {
	return roleNames.every(role => userRoles.some(userRole => userRole.name === role))
}

export const hasSomeRole = (userRoles: Role[], roleNames: string[]): boolean => {
	return roleNames.some(role => userRoles.some(userRole => userRole.name === role))
}
