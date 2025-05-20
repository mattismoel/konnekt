import { requestAndParse } from "@/lib/api";
import { createUrl } from "@/lib/url";
import { z } from "zod";
import { memberSchema, uploadMemberProfilePicture } from "./member";
import { env } from "../../env";

const MINIMUM_PASSWORD_LENGTH = 8
const MAXIMUM_PASSWORD_LENGTH = 24

const baseRegisterForm = z.object({
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
})

export const registerForm = baseRegisterForm.extend({
	profilePictureFile: z
		.instanceof(File)
})
	.refine(({ password, passwordConfirm }) => passwordConfirm === password, {
		message: "Adgangskoder skal være ens",
	})

export type RegisterFormValues = z.infer<typeof registerForm>


export const loginForm = z.object({
	email: z
		.string()
		.email(),
	password: z
		.string()
})

export type LoginFormValues = z.infer<typeof loginForm>

export const login = async (form: LoginFormValues) => {
	const member = await requestAndParse(
		createUrl(`${env.VITE_SERVER_ORIGIN}/api/auth/login`),
		memberSchema,
		"Could not login member",
		{ body: form, bodySchema: loginForm },
		"POST",
	)

	return member
}

export const logOut = async () => {
	await requestAndParse(
		createUrl(`/api/auth/log-out`),
		undefined,
		"Could not log out user",
		undefined,
		"POST"
	)
}

const registerSchema = baseRegisterForm.extend({
	profilePictureUrl: z.string().url()
})

export const register = async (form: RegisterFormValues) => {
	const { profilePictureFile } = form;

	const profilePictureUrl = await uploadMemberProfilePicture(profilePictureFile)

	const member = await requestAndParse(
		createUrl(`/api/auth/register`),
		memberSchema,
		"Could not register member",
		{
			body: {
				...form,
				profilePictureUrl,
			}, bodySchema: registerSchema
		},
		"POST",
	)

	return member
}
