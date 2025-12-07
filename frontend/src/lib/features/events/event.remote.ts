import { form, getRequestEvent, query } from "$app/server";
import { id } from "$lib/model";
import { createFileUrl, createListResult, type PBConcert, type PBEvent } from "$lib/pocketbase";
import { queryOptions } from "$lib/query";
import { parse, startOfToday } from "date-fns";
import { getArtist } from "../artists/artist.remote";
import { concertSchema, eventSchema, type Event } from "./event";
import { getVenue } from "../venues/venue.remote";
import { INPUT_DATETIME_FORMAT } from "$lib/time";
import { redirect } from "@sveltejs/kit";
import z from "zod";

const concertForm = z.object({
	artistId: id,
	fromDate: z.string(),
	toDate: z.string()
});

const createConcertForm = concertForm;

const editConcertForm = concertForm.extend({
	id: id.optional()
});

const eventForm = z.object({
	title: z.string().nonempty(),
	description: z.string().nonempty(),
	ticketUrl: z.url(),
	venueId: id,
	isPublic: z.coerce.boolean<boolean>()
});

const createEventForm = eventForm.extend({
	concerts: createConcertForm.array().min(1),
	cover: z.file().max(5_000_000)
});

const editEventForm = eventForm.extend({
	eventId: id,
	concerts: editConcertForm.array().min(1),
	cover: z.file().max(5_000_000).optional()
});

export type ConcertForm = z.infer<typeof concertForm>;

export const getUpcomingEvents = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const filterStr = locals.pb.filter("concerts.fromDate >= {:today}", {
		today: startOfToday().toISOString()
	});

	const result = await getEvents({
		filter: opts?.filter ? [filterStr, opts?.filter].join("&&") : filterStr
	});

	return result;
});

export const getEvents = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const { items, ...rest } = await locals.pb
		.collection("events")
		.getList<PBEvent>(opts?.page, opts?.perPage, {
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
	const { locals } = getRequestEvent();

	const concertBatch = locals.pb.createBatch();
	data.concerts.forEach(async (concert) => {
		const data = {
			fromDate: parse(concert.fromDate, INPUT_DATETIME_FORMAT, new Date()),
			toDate: parse(concert.toDate, INPUT_DATETIME_FORMAT, new Date()),
			artist: concert.artistId
		};

		concertBatch.collection("concerts").create(data);
	});

	const results = await concertBatch.send();
	const concertIds = results.map((result) => result.body.id as string);
	const postData = { ...data, venue: data.venueId, concerts: concertIds };

	const { id } = await locals.pb.collection("events").create(postData);

	redirect(302, `/events/${id}`);
});

export const editEvent = form(editEventForm, async (data) => {
	const { locals } = getRequestEvent();

	let concertIds = data.concerts.flatMap((c) => (c.id ? c.id : []));

	await Promise.all(
		data.concerts.map(async (concert) => {
			const data = {
				fromDate: parse(concert.fromDate, INPUT_DATETIME_FORMAT, new Date()),
				toDate: parse(concert.toDate, INPUT_DATETIME_FORMAT, new Date()),
				artist: concert.artistId
			};

			// If concert already exists, we want to update it, else create new concert and add ID.
			if (concert.id) {
				await locals.pb.collection("concerts").update(concert.id, data);
			} else {
				const { id } = await locals.pb.collection("concerts").create(data);
				concertIds.push(id);
			}
		})
	);

	const postData = { ...data, venue: data.venueId, concerts: concertIds };

	await locals.pb.collection("events").update(data.eventId, postData);
	redirect(302, `/events/${data.eventId}`);
});
