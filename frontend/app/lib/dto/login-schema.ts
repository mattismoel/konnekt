import { z } from "zod";

export const loginSchema = z.object({
  email: z.string().trim().email(),
  password: z.string().trim().min(1)
})

export type LoginLoad = z.infer<typeof loginSchema>
