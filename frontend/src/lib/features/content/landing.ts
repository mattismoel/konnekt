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
		landingImageSchema.array(),
	)

	return srcs
}

export const uploadLandingImage = async (file: File): Promise<Image> => {
	const formData = new FormData()

	formData.set("file", file)

	const res = await fetch("/api/content/landing-images", {
		body: formData,
		method: "POST",
		credentials: "include",
	})

	if (!res.ok) {
		const error = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not post new landing image", error.message)
	}

	const img = landingImageSchema.parse(await res.json())

	return img
}

export const deleteLandingImage = async (id: ID) => {
	await requestAndParse(
		createUrl(`/api/content/landing-images/${id}`),
		undefined,
		"Could not delete landing page image",
		undefined,
		"DELETE"
	)
}
