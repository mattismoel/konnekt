import { form, getRequestEvent, query } from "$app/server";
import { id } from "$lib/model";
import { createFileUrl, type PBArtist, type PBGenre } from "$lib/pocketbase";
import { queryOptions } from "$lib/query";
import z from "zod";
import { artistSchema, genreSchema } from "./artist";
import { getPreviousEvents, getUpcomingEvents } from "../events/event.remote";

const genreForm = z.object({
	name: z
		.string()
		.nonempty()
		.transform((v) => v.charAt(0).toUpperCase() + v.slice(1))
});

const artistForm = z.object({
	name: z.string().nonempty(),
	description: z.string().nonempty(),
	genreIds: id.array().min(1),
	previewUrl: z.url().optional()
});

const createArtistForm = artistForm.extend({
	cover: z.file().max(5_000_000).mime(["image/png", "image/jpeg", "image/webp"]),
	socials: z
		.object({
			url: z.url()
		})
		.array()
		.optional()
});

const editArtistForm = artistForm.extend({
	artistId: id,
	cover: z.file().max(5_000_000).mime(["image/png", "image/jpeg", "image/webp"]).optional(),
	socials: z
		.object({
			url: z.url(),
			id: id.optional()
		})
		.array()
		.optional()
});

export const getArtist = query(id, async (artistId) => {
	const { locals } = getRequestEvent();

	const record = await locals.pb
		.collection("artists")
		.getOne<PBArtist>(artistId, { expand: "genres,socials" });

	const genres = record.expand?.genres || [];
	const socials = record.expand?.socials;

	const artist = artistSchema.parse({
		...record,
		socials: socials,
		genres: genres,
		cover: createFileUrl("artists", record.id, record.cover)
	});

	return artist;
});

export const getArtists = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const { items, ...rest } = await locals.pb
		.collection("artists")
		.getList<PBArtist>(opts?.page, opts?.perPage, {
			...opts,
			expand: "socials,genres",
			sort: "name"
		});

	const artists = artistSchema.array().parse(
		items.map((record) => ({
			...record,
			cover: createFileUrl("artists", record.id, record.cover),
			genres: record.expand?.genres ?? [],
			socials: record.expand?.socials
		}))
	);

	return { items: artists, ...rest };
});

export const createArtist = form(createArtistForm, async (data) => {
	const { locals } = getRequestEvent();

	const socialIds: string[] = [];
	if (data.socials) {
		await Promise.all(
			data.socials.map(async (social) => {
				const { id } = await locals.pb.collection("socials").create({ url: social.url });
				socialIds.push(id);
			})
		);
	}

	await locals.pb.collection("artists").create({
		...data,
		socials: socialIds,
		genres: data.genreIds
	});
});

export const editArtist = form(editArtistForm, async (data) => {
	const { locals } = getRequestEvent();

	let socialIds = data.socials?.flatMap((social) => social.id ?? []);

	if (data.socials) {
		await Promise.all(
			data.socials?.map(async (social) => {
				if (social.id) {
					await locals.pb.collection("socials").update(social.id, { url: social.url });
					return;
				}

				const { id } = await locals.pb.collection("socials").create({ url: social.url });
				socialIds?.push(id);
			})
		);
	}

	await locals.pb.collection("artists").update(data.artistId, {
		...data,
		genres: data.genreIds,
		socials: socialIds
	});
});

export const getUpcomingArtists = query(async () => {
	const { items: upcomingEvents } = await getUpcomingEvents(undefined);

	const artists = upcomingEvents.flatMap((event) =>
		event.concerts.flatMap((concert) => concert.artist)
	);

	return artists;
});

export const getPreviousArtists = query(async () => {
	const { items: previousEvents } = await getPreviousEvents(undefined);

	const artists = previousEvents.flatMap((event) =>
		event.concerts.flatMap((concert) => concert.artist)
	);

	return artists;
});

export const getGenres = query(async () => {
	const { locals } = getRequestEvent();

	const records = await locals.pb.collection("genres").getFullList<PBGenre>();

	const genres = genreSchema.array().parse(records);
	return genres;
});

export const createGenre = form(genreForm, async (data) => {
	const { locals } = getRequestEvent();
	await locals.pb.collection("genres").create(data);
});

export const deleteGenre = form(z.object({ id: id }), async (data) => {
	const { locals } = getRequestEvent();
	await locals.pb.collection("genres").delete(data.id);
});
