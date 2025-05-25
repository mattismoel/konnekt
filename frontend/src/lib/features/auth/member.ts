import { APIError, apiErrorSchema, idSchema, requestAndParse, type ID } from "@/lib/api"
import { createListResult } from "@/lib/query"
import { createUrl, type Query } from "@/lib/url"
import { z } from "zod"
import { setMemberTeams, teamSchema } from "./team"

export const memberSchema = z.object({
	id: idSchema,
	email: z.string().email(),
	firstName: z.string(),
	lastName: z.string(),
	teams: teamSchema.array().min(1),
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
	memberTeams: z
		.number()
		.int()
		.positive()
		.array(),
	image: z.instanceof(File).optional()
})

export type MemberFormValues = z.infer<typeof memberForm>

const editMemberSchema = memberForm
	.omit({ image: true })
	.extend({ profilePictureUrl: z.string().url().optional() })

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

export const approveMember = async (memberId: ID) => {
	return requestAndParse(
		createUrl(`/api/members/${memberId}/approve`),
		undefined,
		"Could not approve member",
		undefined,
		"POST"
	)
}

export const deleteMember = async (memberId: ID) => {
	return requestAndParse(
		createUrl(`/api/members/${memberId}`),
		undefined,
		"Could not delete member",
		undefined,
		"DELETE"
	)
}

export const memberById = async (memberId: ID) => {
	const member = await requestAndParse(
		createUrl(`/api/members/${memberId}`),
		memberSchema,
		"Could not get member by ID"
	)

	return member
}

export const editMember = async (memberId: ID, form: MemberFormValues) => {
	const { data, success, error } = memberForm.safeParse(form)
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

	await setMemberTeams(memberId, form.memberTeams)

	return member
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
