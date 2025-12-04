import { form, getRequestEvent, query } from "$app/server";
import { id } from "$lib/model";
import { createFileUrl, createListResult, type PBConcert, type PBEvent } from "$lib/pocketbase";
import { queryOptions } from "$lib/query";
import { startOfToday } from "date-fns";
import { getArtist } from "./artist.remote";
import { concertSchema, eventSchema, type Event } from "./event";
import { getVenue } from "./venue.remote";
import { z } from "zod";

const concertForm = z.object({
	artistId: id,
	fromDate: z.string(),
	toDate: z.string()
});

const eventForm = z.object({
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.url(),
	venueId: id,
	concerts: concertForm.array().min(1),
	isPublic: z.coerce.boolean<boolean>()
});

const createEventForm = eventForm.extend({
	cover: z.file().max(5_000_000)
});

const editEventForm = eventForm
	.extend({
		cover: z.union([z.file().max(5_000_000), z.undefined()])
	})
	.partial();

export type ConcertForm = z.infer<typeof concertForm>;

export const getUpcomingEvents = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const { items, ...rest } = await locals.pb
		.collection("events")
		.getList<PBEvent>(opts?.page, opts?.perPage, {
			filter: locals.pb.filter("concerts.fromDate > {:date}", {
				date: startOfToday().toISOString()
			}),
			...opts
		});

	let events: Event[] = [];

	await Promise.all(
		items.map(async (item, i) => {
			const venue = await getVenue(item.venue);
			const concerts = await getEventConcerts(item.id);
			events[i] = { ...item, concerts, venue, cover: createFileUrl("events", item.id, item.cover) };
		})
	);

	const result = createListResult(eventSchema).parse({ ...rest, items: events });
	return result;
});

export const getPreviousEvents = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const { items, ...rest } = await locals.pb
		.collection("events")
		.getList<PBEvent>(opts?.page, opts?.perPage, {
			filter: locals.pb.filter("concerts.fromDate <= {:date}", {
				date: startOfToday().toISOString()
			}),
			...opts
		});

	let events: Event[] = [];

	await Promise.all(
		items.map(async (item, i) => {
			const venue = await getVenue(item.venue);
			const concerts = await getEventConcerts(item.id);
			events[i] = { ...item, concerts, venue, cover: createFileUrl("events", item.id, item.cover) };
		})
	);

	const result = createListResult(eventSchema).parse({ ...rest, items: events });
	return result;
});

export const getEventConcerts = query(id, async (eventId) => {
	const { locals } = getRequestEvent();

	const { concerts: concertIds } = await locals.pb
		.collection("events")
		.getOne<Pick<PBEvent, "concerts">>(eventId, {
			fields: "concerts"
		});

	const concerts = await Promise.all(
		concertIds.map(async (id) => {
			const { artist: artistId, ...concert } = await locals.pb
				.collection("concerts")
				.getOne<PBConcert>(id);

			const artist = await getArtist(artistId);

			return { ...concert, artist };
		})
	);

	return concertSchema.array().parse(concerts);
});

export const getEvent = query(id, async (eventId) => {
	const { locals } = getRequestEvent();
	const record = await locals.pb.collection("events").getOne<PBEvent>(eventId);

	const concerts = await getEventConcerts(eventId);

	const venue = await getVenue(record.venue);

	const event = eventSchema.parse({
		...record,
		venue,
		concerts,
		cover: createFileUrl("events", record.id, record.cover)
	});

	return event;
});

export const createEvent = form(createEventForm, async (data) => {
	console.log("Hello");
	console.log(data);
});

export const editEvent = form(editEventForm, async (data) => {});
