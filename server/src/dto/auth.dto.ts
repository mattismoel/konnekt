import { z } from "zod"
import { MAX_PASSWORD_LENGTH, MIN_PASSWORD_LENGTH } from "@/shared/auth/constant"

export const registerSchema = z.object({
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
  password: z
    .string()
    .trim()
    .min(MIN_PASSWORD_LENGTH, { message: `Password must be at least ${MIN_PASSWORD_LENGTH} characters` })
    .max(MAX_PASSWORD_LENGTH, { message: `Password must be a max of ${MAX_PASSWORD_LENGTH} characters` }),
  passwordConfirm: z
    .string()
})
  .refine(({ password, passwordConfirm }) => password == passwordConfirm, { message: "Passwords do not match" })

export const loginSchema = z.object({
  email: z
    .string({ message: "Email must be set" })
    .trim()
    .email(),
  password: z
    .string({ message: "Password must be set" })
    .trim()
})
export type RegisterDTO = z.infer<typeof registerSchema>
export type LoginDTO = z.infer<typeof loginSchema>
