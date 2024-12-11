import { z } from "zod";

export const venueSchema = z.object({
  name: z.string(),
  country: z.string(),
  city: z.string(),
})

export type AddressDTO = z.infer<typeof venueSchema>
