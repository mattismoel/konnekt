import { z } from "zod";

const envSchema = z.object({
  MODE: z.union([
    z.literal("development"),
    z.literal("production"),
  ])
    .default("development"),
  BASE_URL: z.union([z.string().url(), z.literal("/")]),
  PROD: z.boolean(),
  DEV: z.boolean(),
  SSR: z.boolean(),
})

export const env = () => envSchema.parse(import.meta.env)
