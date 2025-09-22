import z from "zod";
import { throwAPIError } from "@/lib/api/error";
import { createListResult } from "@/lib/api/list";
import { teamSchema } from "./team";

export const permissionSchema = z.object({
	id: z.int().positive(),
	name: z.string().nonempty(),
	displayName: z.string().nonempty()
})

export type Team = z.infer<typeof teamSchema>
export type Permission = z.infer<typeof permissionSchema>

export const registerFormSchema = z.object({
	email: z.email().nonempty(),
	firstName: z.string().nonempty(),
	lastName: z.string().nonempty(),
	password: z.string().nonempty(),
	passwordConfirm: z.string().nonempty(),
	avatarFile: z.file().mime(["image/png", "image/jpeg"])
})

export const loginFormSchema = z.object({
	email: z.email(),
	password: z.string()
})


export const memberPermissions = async (memberId: number) => {
	const res = await fetch(`/api/members/${memberId}/permissions`)
	if (!res.ok) {
		throwAPIError(await res.json())
	}

	return createListResult(permissionSchema).parse(await res.json())
}

export const memberTeams = async (memberId: number) => {
	const res = await fetch(`/api/members/${memberId}/teams`)
	if (!res.ok) {
		throwAPIError(await res.json())
	}

	return createListResult(teamSchema).parse(await res.json())
}
