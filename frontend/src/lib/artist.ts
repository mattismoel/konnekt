import { z } from "zod";
import { genreSchema } from "./genre";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { createListResult, type ListResult } from "./list-result";
import { APIError, apiErrorSchema } from "./error";

export const artistSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	imageUrl: z.string().url(),
	description: z.string(),
	genres: genreSchema.array(),
	socials: z.string().url().array()
})

export type Artist = z.infer<typeof artistSchema>

export const artistFormSchema = z.object({
	name: z
		.string()
		.nonempty({ message: "Kustnernavn skal være defineret" }),
	description: z
		.string()
		.nonempty({ message: "Kunstnerbeskreivelse skal være defineret" }),
	image: z
		.instanceof(File)
		.nullable(),
	genreIds: z.number()
		.positive()
		.array()
		.min(1, { message: "Mindst én genre skal være valgt" }),
	socials: z
		.string()
		.nonempty()
		.url({ message: "URL skal være gyldigt" })
		.array(),
})

export type ArtistForm = z.infer<typeof artistFormSchema>
/**
/**
 * @description Lists artists as a {ListResult} object.
 */
export const listArtists = async (params?: URLSearchParams): Promise<ListResult<Artist>> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/artists?` + (params || ""))

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not list artists", err.message)
	}

	const result = createListResult(artistSchema).parse(await res.json())

	return result
}
