import { requestAndParse } from "$lib/api";
import { createUrl } from "$lib/url";
import { z } from "zod";
import { memberSchema, uploadMemberProfilePicture } from "./member";

const MINIMUM_PASSWORD_LENGTH = 8
const MAXIMUM_PASSWORD_LENGTH = 24

export const registerForm = z.object({
	email: z
		.string()
		.email(),
	firstName: z
		.string()
		.nonempty(),
	lastName: z
		.string()
		.nonempty(),
	password: z
		.string()
		.min(MINIMUM_PASSWORD_LENGTH)
		.max(MAXIMUM_PASSWORD_LENGTH),
	passwordConfirm:
		z.string(),
	profilePictureUrl: z
		.string()
		.url()
		.optional()
})
	.refine(({ password, passwordConfirm }) => passwordConfirm === password, {
		message: "Adgangskoder skal være ens",
	})

export const loginForm = z.object({
	email: z
		.string()
		.email(),
	password: z
		.string()
})

export const login = async (fetchFn: typeof fetch, form: z.infer<typeof loginForm>) => {
	const member = await requestAndParse(
		fetchFn,
		createUrl(`/api/auth/login`),
		memberSchema,
		"Could not login member",
		{ body: form, bodySchema: loginForm },
		"POST",
	)

	return member
}

export const logOut = async (fetchFn: typeof fetch) => {
	await requestAndParse(
		fetchFn,
		createUrl(`/api/auth/log-out`),
		undefined,
		"Could not log out user",
		undefined,
		"POST"
	)
}

export const register = async (fetchFn: typeof fetch, form: z.infer<typeof registerForm>, profilePictureFile?: File | null) => {
	let profilePictureUrl: string | undefined

	if (profilePictureFile) {
		profilePictureUrl = await uploadMemberProfilePicture(fetchFn, profilePictureFile)
	}

	const member = await requestAndParse(
		fetchFn,
		createUrl(`/api/auth/register`),
		memberSchema,
		"Could not register member",
		{
			body: {
				...form,
				...(profilePictureUrl && { profilePictureUrl }),
			}, bodySchema: registerForm
		},
		"POST",
	)

	return member
}
