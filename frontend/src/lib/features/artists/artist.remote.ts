import { form, getRequestEvent, query } from "$app/server";
import { id } from "$lib/model";
import { createFileUrl, type PBArtist, type PBGenre } from "$lib/pocketbase";
import { queryOptions } from "$lib/query";
import z from "zod";
import { artistSchema, genreSchema } from "./artist";

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
