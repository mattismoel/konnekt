import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { requestAndParse } from "$lib/api"
import { createListResult } from "$lib/query"
import { createUrl, type Query } from "$lib/url"
import { z } from "zod"
import { permissionSchema } from "./permission"
import { teamSchema } from "./team"

export const memberSchema = z.object({
	id: z.number().positive(),
	email: z.string().email(),
	firstName: z.string(),
	lastName: z.string(),

	profilePictureUrl: z
		.string()
		.url()
		.optional(),

	active: z.boolean(),

	teams: teamSchema.array(),
	permissions: permissionSchema.array(),
})

export type Member = z.infer<typeof memberSchema>

export const editMemberForm = z.object({
	firstName: z
		.string()
		.nonempty(),
	lastName: z
		.string()
		.nonempty(),
	email: z
		.string()
		.email(),
})

export const setMemberTeamsForm = z
	.number()
	.int()
	.positive()
	.array()
	.nonempty()

export const memberSession = async (fetchFn: typeof fetch) => {
	const member = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/session`),
		memberSchema,
		"Could not fetch member session",
	)

	return member
}

export const listMembers = async (fetchFn: typeof fetch, query?: Query) => {
	const result = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/members`, query),
		createListResult(memberSchema),
		"Could not fetch members",
	)

	return result
}

export const approveMember = async (fetchFn: typeof fetch, memberId: number) => {
	return requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/members/${memberId}/approve`),
		undefined,
		"Could not approve member",
		undefined,
		"POST"
	)
}

export const deleteMember = async (fetchFn: typeof fetch, memberId: number) => {
	return requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/members/${memberId}`),
		undefined,
		"Could not delete member",
		undefined,
		"DELETE"
	)
}

export const memberById = async (fetchFn: typeof fetch, memberId: number) => {
	const member = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/members/${memberId}`),
		memberSchema,
		"Could not get member by ID"
	)

	return member
}

export const editMember = async (fetchFn: typeof fetch, memberId: number, form: z.infer<typeof editMemberForm>) => {
	await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/members/${memberId}`),
		undefined,
		"Could not edit member",
		{ bodySchema: editMemberForm, body: form },
		"PUT"
	)
}

export const setMemberTeams = async (fetchFn: typeof fetch, memberId: number, teamIds: number[]) => {
	await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/members/${memberId}/teams`),
		undefined,
		"Could not set member teams",
		{ bodySchema: setMemberTeamsForm, body: teamIds },
		"PUT"
	)
}
