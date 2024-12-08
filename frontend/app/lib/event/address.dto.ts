import { z } from "zod";

export const addressSchema = z.object({
  country: z.string(),
  city: z.string(),
  street: z.string(),
  houseNumber: z.string()
})

export type AddressDTO = z.infer<typeof addressSchema>
