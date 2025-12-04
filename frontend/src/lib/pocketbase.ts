import { PUBLIC_POCKETBASE_URL } from "$env/static/public";
import type { RecordModel } from "pocketbase";
import { z, ZodType } from "zod";

const pbId = z.string().nonempty();

const baseFields = z.object({
	id: pbId,
	created: z.coerce.date(),
	updated: z.coerce.date()
});

export const pbVenue = z.object({
	...baseFields.shape,
	name: z.string().nonempty(),
	city: z.string().nonempty(),
	countryCode: z.string().nonempty()
});

export const pbSocial = z.object({
	...baseFields.shape,
	url: z.url()
});

export const pbGenre = z.object({
	...baseFields.shape,
	name: z.string().nonempty()
});

export const pbArtist = z.object({
	...baseFields.shape,
	name: z.string().nonempty(),
	description: z.string().nonempty(),
	cover: z.string().nonempty(),
	genres: pbId.array().min(1),
	socials: pbId.array(),
	previewUrl: z.union([z.url(), z.null()]).optional(),
	expand: z
		.object({
			genres: pbGenre.array().min(1),
			socials: pbSocial.array()
		})
		.optional()
});

export const pbConcert = z.object({
	...baseFields.shape,
	fromDate: z.coerce.date(),
	toDate: z.coerce.date(),
	artist: pbId,
	expand: z
		.object({
			artist: pbArtist
		})
		.optional()
});

export const pbEvent = z.object({
	...baseFields.shape,
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	concerts: pbId.array().min(1),
	venue: pbId,
	ticketUrl: z.url(),
	cover: z.string(),
	expand: z
		.object({
			concert: pbConcert.array().min(1),
			venue: pbVenue
		})
		.optional()
});

export const createFileUrl = (
	collection: string,
	recordId: string,
	path: string,
	thumb?: string
) => {
	const url = new URL(`/api/files/${collection}/${recordId}/${path}`, PUBLIC_POCKETBASE_URL);
	if (thumb) {
		url.searchParams.set("thumb", thumb);
	}

	return url.toString();
};

export const createListResult = <T extends ZodType>(schema: T) =>
	z.object({
		items: schema.array(),
		page: z.int().positive(),
		perPage: z.int().positive(),
		totalItems: z.int().nonnegative(),
		totalPages: z.int().nonnegative()
	});

export type PBEvent = z.infer<typeof pbEvent>;
export type PBConcert = z.infer<typeof pbConcert>;
export type PBVenue = z.infer<typeof pbVenue>;

export type PBArtist = z.infer<typeof pbArtist>;
export type PBGenre = z.infer<typeof pbGenre>;
export type PBSocial = z.infer<typeof pbSocial>;
