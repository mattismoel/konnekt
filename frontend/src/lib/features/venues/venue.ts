import z from "zod";

export const venueSchema = z.object({
	id: z.int().positive(),
	name: z.string().nonempty(),
	city: z.string().nonempty(),
	country: z.string().nonempty()
})
