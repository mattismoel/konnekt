import { z } from "zod";

const envSchema = z.object({
  ENV: z.union([
    z.literal("production"),
    z.literal("testing"),
    z.literal("development")
  ])
    .default("development"),
  PORT: z.coerce.number(),
  DSN: z.string(),
})

export const env = envSchema.parse(process.env)
