import { z } from "zod";

export type VenueDTO = {
  id: number;
  name: string;
  country: string;
  city: string;
}

export const createVenueSchema = z.object({
  name: z
    .string()
    .trim()
    .min(1),
  country: z
    .string()
    .trim()
    .min(1),
  city: z
    .string()
    .trim()
    .min(1),
})

export type CreateVenueDTO = z.infer<typeof createVenueSchema>
