import z, { ZodType } from "zod";

export const createListResult = <T extends ZodType>(schema: T) => {
	return z.object({
		records: schema.array()
	})
}
