import { z } from "zod";

export const userSchema = z.object({
  id: z.number().min(1),
  email: z.string().email(),
  firstName: z.string().min(1),
  lastName: z.string().min(1),
  roles: z.union([
    z.literal("admin"),
    z.literal("user"),
    z.literal("guest"),
  ]).array()
})

export type User = z.infer<typeof userSchema>
