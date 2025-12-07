import { form, getRequestEvent, query } from "$app/server";
import { id } from "$lib/model";
import type { PBVenue } from "$lib/pocketbase";
import { queryOptions } from "$lib/query";
import z from "zod";
import { venueSchema } from "./venue";

const venueForm = z.object({
	name: z.string().nonempty(),
	countryCode: z.string().nonempty(),
	city: z.string().nonempty()
});

const createVenueForm = venueForm;
const editVenueForm = venueForm.extend({
	venueId: id
});

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

export const createVenue = form(createVenueForm, async (data) => {
	const { locals } = getRequestEvent();
	await locals.pb.collection("venues").create(data);
});

export const editVenue = form(editVenueForm, async ({ venueId, ...rest }) => {
	const { locals } = getRequestEvent();
	await locals.pb.collection("venues").update(venueId, rest);
});
