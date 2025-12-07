import { baseFields } from "$lib/model";
import {} from "@sveltejs/kit";
import z from "zod";

const MINIMUM_PASSWORD_LENGTH = 8;
const MAXIMUM_PASSWORD_LENGTH = 24;

export const teamSchema = z.object({});

export const userSchema = z.object({
	...baseFields.shape,
	email: z.email().optional(),
	firstName: z.string().nonempty(),
	lastName: z.string().nonempty(),
	avatar: z.union([z.url(), z.literal("")]),
	teams: teamSchema.array().min(1),
	approved: z.boolean()
});

export const registerForm = z
	.object({
		email: z.email(),
		firstName: z.string().nonempty(),
		lastName: z.string().nonempty(),
		password: z.string().min(MINIMUM_PASSWORD_LENGTH).max(MAXIMUM_PASSWORD_LENGTH),
		passwordConfirm: z.string(),
		avatar: z.file().max(5_000_000)
	})
	.refine((data) => data.password === data.passwordConfirm, "Adgangskoder er ikke ens");

export const loginForm = z.object({
	email: z.email(),
	password: z.string()
});

export type User = z.infer<typeof userSchema>;

export type RegisterForm = z.infer<typeof registerForm>;
export type LoginForm = z.infer<typeof loginForm>;
