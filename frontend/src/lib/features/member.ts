import { APIError, apiErrorSchema, requestAndParse } from "@/lib/api"
import { createListResult } from "@/lib/query"
import { createUrl, type Query } from "@/lib/url"
import { z } from "zod"

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
})

export type Member = z.infer<typeof memberSchema>

export const memberForm = z.object({
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

export const editMemberForm = memberForm
	.extend({ image: z.instanceof(File).nullable() })

const editMemberSchema = editMemberForm
	.omit({ image: true })
	.extend({ profilePictureUrl: z.string().url().optional() })

export const setMemberTeamsForm = z
	.number()
	.int()
	.positive()
	.array()
	.nonempty()

export const memberSession = async () => {
	const member = await requestAndParse(
		createUrl(`/api/auth/session`),
		memberSchema,
		"Could not fetch member session",
	)

	return member
}

export const listMembers = async (query?: Query) => {
	const result = await requestAndParse(
		createUrl(`/api/members`, query),
		createListResult(memberSchema),
		"Could not fetch members",
	)

	return result
}

export const approveMember = async (memberId: number) => {
	return requestAndParse(
		createUrl(`/api/members/${memberId}/approve`),
		undefined,
		"Could not approve member",
		undefined,
		"POST"
	)
}

export const deleteMember = async (memberId: number) => {
	return requestAndParse(
		createUrl(`/api/members/${memberId}`),
		undefined,
		"Could not delete member",
		undefined,
		"DELETE"
	)
}

export const memberById = async (memberId: number) => {
	const member = await requestAndParse(
		createUrl(`/api/members/${memberId}`),
		memberSchema,
		"Could not get member by ID"
	)

	return member
}

export const editMember = async (memberId: number, form: z.infer<typeof editMemberForm>) => {
	const { data, success, error } = editMemberForm.safeParse(form)
	if (!success) throw error

	const { image, ...rest } = data;

	const profilePictureUrl = image ? await uploadMemberProfilePicture(image) : undefined

	const member = requestAndParse(
		createUrl(`/api/members/${memberId}`),
		memberSchema,
		"Could not update artist",
		{ bodySchema: editMemberSchema, body: { ...rest, profilePictureUrl } },
		"PUT"
	)

	return member
}

export const setMemberTeams = async (memberId: number, teamIds: number[]) => {
	await requestAndParse(
		createUrl(`/api/members/${memberId}/teams`),
		undefined,
		"Could not set member teams",
		{ bodySchema: setMemberTeamsForm, body: teamIds },
		"PUT"
	)
}

export const uploadMemberProfilePicture = async (file: File): Promise<string> => {
	const formData = new FormData()

	formData.append("file", file)

	const res = await fetch(`/api/members/picture`, {
		body: formData,
		method: "POST",
		credentials: "include",
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not upload member profile picture", err.message)
	}

	const url = await res.text()

	return url
}
