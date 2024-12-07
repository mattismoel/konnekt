import { createAddressSchema, type AddressDTO } from "@/address/address.dto";
import { createGenreSchema, type GenreDTO } from "@/genre/genre.dto";
import { z } from "zod";

export type EventDTO = {
  id: number;
  title: string;
  description: string;
  fromDate: Date;
  toDate: Date;
  address: AddressDTO;
  genres: string[];
}

export const createEventSchema = z.object({
  title: z.string({ message: "Title is required" }).trim().min(1),
  description: z.string({ message: "Description is required" }).trim().min(1),
  fromDate: z.coerce.date({ message: "Invalid date" }),
  toDate: z.coerce.date({ message: "Invalid date" }),
  address: createAddressSchema,
  genres: createGenreSchema.array().min(1, { message: "There must be at least 1 genre" })
})

export type CreateEventDTO = z.infer<typeof createEventSchema>
