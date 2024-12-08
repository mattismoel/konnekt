import { z } from "zod";

export type UserDTO = {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
}

export const createUserSchema = z.object({
  firstName: z
    .string({ message: "First name must be set" })
    .trim()
    .min(1, { message: "First name must not be empty" }),
  lastName: z
    .string({ message: "Last name must be set" })
    .trim()
    .min(1, { message: "Last name must not be empty" }),
  email: z
    .string({ message: "Email must be set" })
    .trim()
    .email(),
  passwordHash: z
    .string()
    .min(1)
})

export type CreateUserDTO = z.infer<typeof createUserSchema>
