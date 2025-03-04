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

export const createArtist = async (form: z.infer<typeof artistFormSchema>, init?: RequestInit) => {
	let res = await fetch(`${PUBLIC_BACKEND_URL}/artists`, {
		...init,
		method: 'POST',
		credentials: 'include',
		body: JSON.stringify(form)
	});

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not create artist", err.message)
	}

	const artist = artistSchema.parse(await res.json())

	if (!form.image) return artist

	const imageUrl = await uploadArtistImage(artist.id, form.image)

	artist.imageUrl = imageUrl

	return artist
};

/**
 * @description Updates an artist.
 * @param {number} artistId - The artist to be updated's ID.
 * @param form - The form data to update the artist with.
 * @param {RequestInit} init - Must be specified with request headers if called from server.
 */
export const updateArtist = async (
	artistId: number,
	form: z.infer<typeof artistFormSchema>,
	init?: RequestInit,
): Promise<Artist> => {
	console.log("UPDATE")
	const res = await fetch(`${PUBLIC_BACKEND_URL}/artists/${artistId}`, {
		...init,
		method: 'PUT',
		credentials: "include",
		body: JSON.stringify(form)
	});

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not update artist", err.message)
	}

	if (!form.image) {
		const artist = await artistById(artistId)
		return artist
	}

	const imageUrl = await uploadArtistImage(artistId, form.image)
	const artist = artistSchema.parse({ ...await res.json(), imageUrl })

	artist.imageUrl = imageUrl

	return artist
};

/**
 * @description Uploads the artist image for the artist specified by its artistId.
 * @param {number} artistId  - The ID of the artist.
 * @param {File} file - The image file to be used as the artist image.
 * @returns {string} The URL of the artist image.
 */
export const uploadArtistImage = async (artistId: number, file: File): Promise<string> => {
	const formData = new FormData();
	formData.append('image', file);

	const res = await fetch(`${PUBLIC_BACKEND_URL}/artists/image/${artistId}`, {
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
export const listArtists = async (params?: URLSearchParams): Promise<ListResult<Artist>> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/artists?` + (params || ""))

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, "Could not list artists", err.message)
	}

	const result = createListResult(artistSchema).parse(await res.json())

	return result
}

/**
 * @description Gets an artists by its ID.
 * @param {number} id - The ID of the artist.
 */
export const artistById = async (id: number): Promise<Artist> => {
	const res = await fetch(`${PUBLIC_BACKEND_URL}/artists/${id}`)

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())
		throw new APIError(res.status, `Could not get artist`, err.message)
	}

	const artist = artistSchema.parse(await res.json())

	return artist
}
