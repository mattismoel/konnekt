import { z, type ZodTypeAny } from "zod";

export const createListResult = <T extends ZodTypeAny>(schema: T) => z.object({
  page: z.number().positive(),
  perPage: z.number().positive(),
  pageCount: z.number().nonnegative(),
  totalCount: z.number().nonnegative(),
  records: schema.array()
})

export type ListResult<T> =
  z.infer<ReturnType<typeof createListResult<z.ZodType<T>>>>;
