import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { requestAndParse } from "$lib/api";
import { createUrl } from "$lib/url";
import { z } from "zod";

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

	z.literal("view:permission"),

	z.literal("view:member"),
	z.literal("edit:member"),
	z.literal("delete:member"),
])

export const permissionSchema = z.object({
	id: z.number().int().positive(),
	name: permissionTypes,
	displayName: z.string().nonempty(),
	description: z.string().nonempty(),
})

export type PermissionType = z.infer<typeof permissionTypes>
export type Permission = z.infer<typeof permissionSchema>

export const userPermissions = async (fetchFn: typeof fetch, userId: number): Promise<Permission[]> => {
	const permissions = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/permissions/${userId}`),
		permissionSchema.array(),
		"Could not fetch user permissions",
	)

	return permissions
}

export const hasPermissions = (userPermissions: Permission[], permNames: PermissionType[]): boolean => {
	return permNames.every(perm => userPermissions.some(userPerm => userPerm.name === perm))
}
