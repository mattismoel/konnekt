import { z } from "zod"

export const SESSION_COOKIE_NAME = "konnekt-session"
export const SESSION_COOKIE_LIFETIME_MILLIS = 30 * 60000 * 60 * 24 // 30 days.

export const roleSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	description: z.string()
})
