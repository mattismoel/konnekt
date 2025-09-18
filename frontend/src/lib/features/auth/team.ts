import z from "zod";

export const team = z.object({
	id: z.int().positive(),
	name: z.string().nonempty(),
	displayName: z.string().nonempty(),
})
