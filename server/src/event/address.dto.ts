import { z } from "zod";

export const createAddressSchema = z.object({
  country: z
    .string({ message: "Country is required" })
    .trim()
    .min(1, { message: "Country must not be empty" }),
  city: z
    .string({ message: "City is required" })
    .trim()
    .min(1, { message: "City must not be empty" }),
  street: z
    .string({ message: "Street is required" })
    .trim()
    .min(1, ". must not be empty"),
  houseNumber: z
    .string({ message: "House number is required" })
    .trim()
    .min(1, { message: "House must not be empty" }),
})

export type CreateAddressDTO = z.infer<typeof createAddressSchema>

export type AddressDTO = {
  id: number;
  country: string;
  city: string;
  street: string;
  houseNumber: string;
}
