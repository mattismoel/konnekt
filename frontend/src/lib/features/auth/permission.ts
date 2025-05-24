import { idSchema, requestAndParse, type ID } from "@/lib/api";
import { createUrl } from "@/lib/url";
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

	z.literal("view:team"),
	z.literal("edit:team"),
	z.literal("delete:team"),

	z.literal("view:genre"),
	z.literal("edit:genre"),
	z.literal("delete:genre"),

	z.literal("view:permission"),

	z.literal("view:member"),
	z.literal("edit:member"),
	z.literal("delete:member"),

	z.literal("edit:content"),
	z.literal("delete:content"),
])

export const permissionSchema = z.object({
	id: idSchema,
	name: permissionTypes,
	displayName: z.string().nonempty(),
	description: z.string().nonempty(),
})

export type PermissionType = z.infer<typeof permissionTypes>
export type Permission = z.infer<typeof permissionSchema>

export const memberPermissions = async (memberId: ID): Promise<Permission[]> => {
	const permissions = await requestAndParse(
		createUrl(`/api/members/${memberId}/permissions`),
		permissionSchema.array(),
		"Could not fetch member permissions",
	)

	return permissions
}
