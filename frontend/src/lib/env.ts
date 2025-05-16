import { z } from "zod";

const envSchema = z.object({
  MODE: z.union([
    z.literal("development"),
    z.literal("production"),
  ])
    .default("development"),
  // BASE_URL: z.string().url()
  PROD: z.boolean(),
  DEV: z.boolean(),
  SSR: z.boolean(),
  VITE_SERVER_ORIGIN: z.string().url()
})

type Env = z.infer<typeof envSchema>

export const env = envSchema.parse(import.meta.env)
