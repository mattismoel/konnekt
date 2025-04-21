import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { requestAndParse } from "$lib/api";
import { createUrl } from "$lib/url";
import { z } from "zod";
import { memberSchema } from "./member";

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
		z.string()
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
	const user = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/login`),
		memberSchema,
		"Could not login member",
		{ body: form, bodySchema: loginForm },
		"POST",
	)

	return user
}

export const register = async (fetchFn: typeof fetch, form: z.infer<typeof registerForm>) => {
	const user = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/auth/register`),
		memberSchema,
		"Could not register member",
		{ body: form, bodySchema: registerForm },
		"POST",
	)

	return user
}
