import { APIError, apiErrorSchema, idSchema, requestAndParse, type ID } from "@/lib/api"
import { createUrl } from "@/lib/url"
import { z } from "zod"

export const landingImageSchema = z.object({
	id: idSchema,
	url: z.string().url()
})

export type Image = z.infer<typeof landingImageSchema>

export const landingImages = async () => {
	const srcs = await requestAndParse(
		createUrl("/api/content/landing-images"),
		landingImageSchema.array().min(1),
	)

	return srcs
}

