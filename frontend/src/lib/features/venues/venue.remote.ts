import { getRequestEvent, query } from "$app/server";
import { id } from "$lib/model";
import type { PBVenue } from "$lib/pocketbase";
import { queryOptions } from "$lib/query";
import { venueSchema } from "./venue";

export const getVenue = query(id, async (id) => {
	const { locals } = getRequestEvent();
	const record = await locals.pb.collection("venues").getOne(id);

	return venueSchema.parse(record);
});

export const getVenues = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const { items, ...rest } = await locals.pb
		.collection("venues")
		.getList<PBVenue>(opts?.page, opts?.perPage);

	const venues = venueSchema.array().parse(items);

	return { items: venues, ...rest };
});
