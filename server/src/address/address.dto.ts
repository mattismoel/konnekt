import { z } from "zod";

export type AddressDTO = {
  id: number;
  country: string;
  city: string;
  street: string;
  houseNumber: string;
}

export const createAddressSchema = z.object({
  country: z.string().trim().min(1),
  city: z.string().trim().min(1),
  street: z.string().trim().min(1),
  houseNumber: z.string().trim().min(1),
})

export type CreateAddressDTO = z.infer<typeof createAddressSchema>

