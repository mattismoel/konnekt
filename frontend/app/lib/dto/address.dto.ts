import { z } from "zod";

export const venueSchema = z.object({
  id: z.number(),
  name: z.string(),
  country: z.string(),
  city: z.string(),
})

export type VenueDTO = z.infer<typeof venueSchema>

export const createVenueSchema = z.object({
  name: z.string().trim().min(1),
  country: z.string().trim().min(1),
  city: z.string().trim().min(1),
})

export type CreateVenueDTO = z.infer<typeof createVenueSchema>
