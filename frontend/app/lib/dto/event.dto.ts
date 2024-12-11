import { z } from "zod";
import { venueSchema } from "./address.dto";

export const eventSchema = z.object({
  id: z.number(),
  title: z.string(),
  description: z.string(),
  coverImageUrl: z.string(),
  fromDate: z.coerce.date(),
  toDate: z.coerce.date(),
  venue: venueSchema,
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
  fromDate: z
    .coerce
    .date(),
  coverImageUrl: z
    .string()
    .url(),
  toDate: z
    .coerce
    .date(),
  city: z
    .string()
    .trim()
    .min(1),
  venue: z
    .string()
    .trim()
    .min(1),
  country: z
    .string()
    .trim()
    .min(1),
  genres: z
    .string()
    .array()
    .min(1, { message: "At least one genre must be set" })
})

export type CreateEditEventDTO = z.infer<typeof createEditEventSchema>
