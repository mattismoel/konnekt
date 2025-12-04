import z from "zod";

export const queryOptions = z.object({
	page: z.int().positive().optional(),
	perPage: z.int().positive().optional(),
	filter: z.string().optional(),
	sort: z.string().optional()
});
