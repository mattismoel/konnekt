import { z } from "zod";
import { genreSchema } from "./genre";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { createListResult, type ListResult } from "./list-result";
import { APIError, apiErrorSchema } from "./error";
import { requestAndParse } from "./api";
import { createUrl, type Query } from "./url";

export const artistSchema = z.object({
	id: z.number().positive(),
	name: z.string(),
	imageUrl: z.string().optional().or(z.string().url().optional()),
	description: z.string(),
	genres: genreSchema.array(),
	socials: z.string().url().array(),
	previewUrl: z.string().url(),
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
	previewUrl: z
		.string()
		.url({ message: "URL skal være gyldigt" })
		.refine(url => {
			let { hostname } = new URL(url);
			hostname = hostname.replace(/^www\./, '');
			return hostname === "open.spotify.com"
		}, { message: "Preview URL skal være fra Spotify" }),
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

const createArtistSchema = artistFormSchema
	.omit({
		image: true
	})
	.extend({
		imageUrl: z.string().url()
	})

const updateArtistSchema = createArtistSchema

export type ArtistForm = z.infer<typeof artistFormSchema>

export const createArtist = async (
	fetchFn: typeof fetch,
	form: z.infer<typeof artistFormSchema>,
) => {
	let { image, ...rest } = form
	if (!image) throw new APIError(400, "Could not upload artist image", "Image file not present")

	const imageUrl = await uploadArtistImage(image)

	const artist = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/artists`),
		artistSchema,
		"Could not create artist",

		{ bodySchema: createArtistSchema, body: { ...rest, imageUrl } },
		"POST",
	)

	return artist
};

/**
 * @description Updates an artist.
 * @param {number} artistId - The artist to be updated's ID.
 * @param form - The form data to update the artist with.
 */
export const updateArtist = async (
	fetchFn: typeof fetch,
	artistId: number,
	form: z.infer<typeof artistFormSchema>,
): Promise<Artist> => {
	const { data, success, error } = artistFormSchema.safeParse(form)
	if (!success) throw error

	const { image, ...rest } = data;

	const imageUrl = image ? await uploadArtistImage(image) : undefined

	const artist = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/artists/${artistId}`),
		artistSchema,
		"Could not update artist",
		{ bodySchema: updateArtistSchema, body: { ...rest, imageUrl } },
		"PUT"
	)

	return artist
};

/**
 * @description Uploads the artist image for the artist specified by its artistId.
 * @param {File} file - The image file to be used as the artist image.
 * @returns {string} The URL of the artist image.
 */
export const uploadArtistImage = async (file: File, init?: RequestInit): Promise<string> => {
	const formData = new FormData();
	formData.append('image', file);

	const res = await fetch(`${PUBLIC_BACKEND_URL}/artists/image`, {
		...init,
		method: 'PUT',
		credentials: 'include',
		body: formData
	});

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not update artist image", err.message)
	}

	const url = await res.text()

	return url
};

/**
 * @description Lists artists as a {ListResult} object.
 */
export const listArtists = async (fetchFn: typeof fetch, query?: Query): Promise<ListResult<Artist>> => {
	const result = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/artists`, query),
		createListResult(artistSchema)
	)

	return result
}

/**
 * @description Gets an artists by its ID.
 * @param {number} id - The ID of the artist.
 */
export const artistById = async (fetchFn: typeof fetch, id: number): Promise<Artist> => {
	const artist = requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/artists/${id}`),
		artistSchema,
	)

	return artist
}


export const deleteArtist = async (fetchFn: typeof fetch, id: number) => {
	await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/artists/${id}`),
		undefined,
		"Could not delete artist",
		undefined,
		"DELETE"
	)
}
