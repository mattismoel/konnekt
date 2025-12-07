import { baseFields } from "$lib/model";
import {} from "@sveltejs/kit";
import z from "zod";

const MINIMUM_PASSWORD_LENGTH = 8;
const MAXIMUM_PASSWORD_LENGTH = 24;

export const teamType = z.union([
	z.literal("admin"),
	z.literal("member"),
	z.literal("event-management"),
	z.literal("public-relations"),
	z.literal("visual-identity"),
	z.literal("project-leader"),
	z.literal("booking"),
	z.literal("economy")
]);

export const permissionType = z.union([
	z.literal("events:view"),
	z.literal("events:edit"),
	z.literal("events:delete"),

	z.literal("artists:view"),
	z.literal("artists:edit"),
	z.literal("artists:delete"),

	z.literal("venues:view"),
	z.literal("venues:edit"),
	z.literal("venues:delete"),

	z.literal("content:view"),
	z.literal("content:edit"),
	z.literal("content:delete")
]);

export const permissionSchema = z.object({
	...baseFields.shape,
	name: permissionType
});

export const teamSchema = z.object({
	...baseFields.shape,
	name: teamType,
	displayName: z.string().nonempty(),
	description: z.string().optional(),
	permissions: permissionSchema.array()
});

export const memberSchema = z.object({
	...baseFields.shape,
	email: z.email().optional(),
	firstName: z.string().nonempty(),
	lastName: z.string().nonempty(),
	avatar: z.union([z.url(), z.literal("")]),
	teams: teamSchema.array().min(1),
	approved: z.boolean(),
	specialRole: z.string().optional()
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

export type Member = z.infer<typeof memberSchema>;
export type TeamType = z.infer<typeof teamType>;
export type Team = z.infer<typeof teamSchema>;

export type RegisterForm = z.infer<typeof registerForm>;
export type LoginForm = z.infer<typeof loginForm>;
