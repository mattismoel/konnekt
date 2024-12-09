import { z } from "zod"

declare global {
  interface Window {
    ENV: z.infer<typeof envSchema>
  }
}

const envSchema = z.object({
  ENV: z
    .union([
      z.literal("development"),
      z.literal("production"),
      z.literal("testing"),
    ])
    .default("development"),
  BACKEND_URL: z
    .string()
    .url()
})

const env = envSchema.parse(process.env)

export default env
