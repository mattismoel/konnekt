import { z } from "zod";

export const landingImage = z.object({
	id: z.string(),
	image: z.string().nonempty()
});
