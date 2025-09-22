import z from "zod";
import { teamSchema } from "./team";
import { throwAPIError } from "@/lib/api/error";

export const memberSchema = z.object({
	id: z.int().positive(),
	email: z.email().nonempty(),
	firstName: z.string().nonempty(),
	lastName: z.string().nonempty(),
	avatarUrl: z.url().nonempty(),
	specialRole: z.string().optional(),
	approved: z.boolean(),
	teams: teamSchema.array(),
})

export type Member = z.infer<typeof memberSchema>
