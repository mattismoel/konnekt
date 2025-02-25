import { z } from "zod";

export const userSchema = z.object({
	id: z.number().positive(),
	email: z.string().email(),
	firstName: z.string(),
	lastName: z.string()
})

export type User = z.infer<typeof userSchema>

