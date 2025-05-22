import { z } from "zod";
import { genreSchema } from "./genre";
import { APIError, apiErrorSchema, idSchema, requestAndParse, type ID } from "@/lib/api";
import { createUrl, isValidUrl, type Query } from "@/lib/url";
import { createListResult, type ListResult } from "@/lib/query";

import type { IconType } from "react-icons/lib";
import {
	FaApple,
	FaExclamationTriangle,
	FaFacebook,
	FaInstagram,
	FaSpotify,
	FaYoutube,
} from "react-icons/fa";

import { env } from "../../env";

export const artistSchema = z.object({
	id: idSchema,
	name: z.string(),
	imageUrl: z.string().optional().or(z.string().url().optional()),
	description: z.string(),
	genres: genreSchema.array(),
	socials: z.string().url().array(),
	previewUrl: z.string().url().optional(),
})

export type Artist = z.infer<typeof artistSchema>

export const artistForm = z.object({
	name: z
		.string()
		.nonempty({ message: "Kustnernavn skal være defineret" }),
	description: z
		.string()
		.nonempty({ message: "Kunstnerbeskreivelse skal være defineret" }),
	previewUrl: z.union([z.literal(""), z
		.string()
		.url({ message: "Spotify preview-URL skal være et gyldigt URL" })
		.refine(url => {
			if (!url) return true
			let { hostname } = new URL(url);
			hostname = hostname.replace(/^www\./, '');
			return hostname === "open.spotify.com"
		}, { message: "Preview URL skal være fra Spotify" }),
	]),
	genreIds: idSchema
		.array()
		.nonempty(),
	socials: z
		.string()
		.nonempty()
		.url({ message: "URL skal være gyldigt" })
		.array(),
	image: z.instanceof(File).optional()
})

export type ArtistFormValues = z.infer<typeof artistForm>

const createArtistSchema = artistForm
	.omit({ image: true })
	.extend({ imageUrl: z.string().url() })

const editArtistSchema = artistForm
	.omit({ image: true })
	.extend({ imageUrl: z.string().url().optional() })

export const createArtist = async (form: ArtistFormValues) => {
	let { image, ...rest } = form
	if (!image) throw new APIError(400, "Could not upload artist image", "Image file not present")

	const imageUrl = await uploadArtistImage(image)

	const artist = requestAndParse(
		createUrl("/api/artists"),
		artistSchema,
		"Could not create artist",

		{ bodySchema: createArtistSchema, body: { ...rest, imageUrl } },
		"POST",
	)

	return artist
};

/**
 * @description Updates an artist.
 * @param {ID} artistId - The artist to be updated's ID.
 * @param form - The form data to update the artist with.
 */
export const updateArtist = async (artistId: ID, form: ArtistFormValues): Promise<Artist> => {
	const { data, success, error } = artistForm.safeParse(form)
	if (!success) throw error

	const { image, ...rest } = data;

	const imageUrl = image ? await uploadArtistImage(image) : undefined

	const artist = requestAndParse(
		createUrl(`/api/artists/${artistId}`),
		artistSchema,
		"Could not update artist",
		{ bodySchema: editArtistSchema, body: { ...rest, imageUrl } },
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

	const res = await fetch("/api/artists/image", {
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
export const listArtists = async (query?: Query): Promise<ListResult<Artist>> => {
	const result = requestAndParse(
		createUrl("/api/artists", {
			orderBy: new Map([["name", "ASC"]]),
			...query,
		}),
		createListResult(artistSchema)
	)

	return result
}

/**
 * @description Gets an artists by its ID.
 * @param {ID} id - The ID of the artist.
 */
export const artistById = async (id: ID): Promise<Artist> => {
	const artist = await requestAndParse(
		createUrl(`/api/artists/${id}`),
		artistSchema,
	)

	return artist
}

export const deleteArtist = async (id: ID) => {
	await requestAndParse(
		createUrl(`/api/artists/${id}`),
		undefined,
		"Could not delete artist",
		undefined,
		"DELETE"
	)
}

const iconMap = new Map<string, IconType>([
	["spotify.com", FaSpotify],
	["open.spotify.com", FaSpotify],
	["instagram.com", FaInstagram],
	["music.apple.com", FaApple],
	["facebook.com", FaFacebook],
	["youtube.com", FaYoutube]
])

export const socialUrlToIcon = (url: string): IconType => {
	if (!isValidUrl(url)) return FaExclamationTriangle

	const { hostname } = new URL(url);
	const iconKey = hostname.replace(/^www\./, '');

	const icon = iconMap.get(iconKey)

	if (!icon) return FaExclamationTriangle

	return icon
}
