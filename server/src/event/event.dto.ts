import { createGenreSchema } from "@/event/genre.dto";
import { z } from "zod";
import { createAddressSchema, type AddressDTO } from "./address.dto";

export type EventDTO = {
  id: number;
  title: string;
  description: string;
  fromDate: Date;
  coverImageUrl: string | null;
  toDate: Date;
  address: AddressDTO;
  genres: string[];
}

export const createEventSchema = z.object({
  title: z
    .string({ message: "Title is required" })
    .trim()
    .min(1, { message: "Title must not be empty" }),
  description: z
    .string({ message: "Description is required" })
    .trim()
    .min(1, { message: "Description must not be empty" }),
  fromDate: z
    .coerce
    .date({ message: "Invalid date" }),
  toDate: z
    .coerce
    .date({ message: "Invalid date" }),
  address: createAddressSchema,
  genres: createGenreSchema.array().min(1, { message: "There must be at least 1 genre" })
})

export type CreateEventDTO = z.infer<typeof createEventSchema>
