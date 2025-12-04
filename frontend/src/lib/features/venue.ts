import { baseFields } from "$lib/model";
import z from "zod";

export const venueSchema = z.object({
	...baseFields.shape,
	name: z.string().nonempty(),
	countryCode: z.string().nonempty(),
	city: z.string().nonempty()
});

export type Venue = z.infer<typeof venueSchema>;
