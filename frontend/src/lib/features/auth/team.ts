import z from "zod";

export const teamSchema = z.object({
	id: z.int().positive(),
	name: z.string().nonempty(),
	displayName: z.string().nonempty(),
})
