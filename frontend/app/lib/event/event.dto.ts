import { z } from "zod";
import { addressSchema } from "./address.dto";

export const eventSchema = z.object({
  id: z.number(),
  title: z.string(),
  description: z.string(),
  coverImageUrl: z.string(),
  fromDate: z.coerce.date(),
  toDate: z.coerce.date(),
  address: addressSchema,
  genres: z.string().array()
})

export type EventDTO = z.infer<typeof eventSchema>
