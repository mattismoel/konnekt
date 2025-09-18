import z from "zod";
import { team } from "./team";

export const member = z.object({
	id: z.int().positive(),
	email: z.email().nonempty(),
	firstName: z.string().nonempty(),
	lastName: z.string().nonempty(),
	avatarUrl: z.url().nonempty(),
	specialRole: z.string().optional(),
	approved: z.boolean(),
	teams: team.array(),
})

export type Member = z.infer<typeof member>
