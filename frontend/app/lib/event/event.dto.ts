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

export const createEditEventSchema = z.object({
  title: z
    .string()
    .trim()
    .min(1),
  description: z
    .string()
    .trim()
    .min(1),
  file: z
    .instanceof(File),
  fromDate: z.coerce.date(),
  toDate: z.coerce.date(),
  city: z.string().trim().min(1),
  venue: z.string().trim().min(1),
  genres: z
    .string()
    .refine(str => str.split(";").length > 0, { message: "At least one genre must be set" })
})

export type CreateEditEventDTO = z.infer<typeof createEditEventSchema>
