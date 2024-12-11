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
    .string({ message: "Titel påkrævet" })
    .trim()
    .min(1, { message: "Titel må ikke være tom" }),
  description: z
    .string({ message: "Beskrivelse påkrævet" })
    .trim()
    .min(1, { message: "Beskrivelse må ikke være tom" }),
  fromDate: z
    .coerce
    .date({ message: "Ugyldig fra-dato" }),
  coverImageUrl: z
    .string({ message: "Cover-billede ugyldigt" })
    .min(1, { message: "Coverbillede skal være sat" })
    .url({ message: "Coverbillede har ugyldigt URL" }),
  toDate: z
    .coerce
    .date({ message: "Ugyldig til-dato" }),
  city: z
    .string({ message: "By er påkrævet" })
    .trim()
    .min(1, { message: "By må ikke være tom" }),
  venue: z
    .string({ message: "Venue er påkrævet" })
    .trim()
    .min(1, { message: "Venue må ikke være tom" }),
  country: z
    .string({ message: "Land er påkrævet" })
    .trim()
    .min(1, { message: "Land må ikke være tomt" }),
  genres: z
    .string({ message: "Genre er ugyldigt format" })
    .refine(str => str.split(";").length > 0, { message: "Mindst én genre påkrævet" })
})

export type CreateEditEventDTO = z.infer<typeof createEditEventSchema>
