<script lang="ts">
	import { DATE_FORMAT } from '$lib/time';

	import { format } from 'date-fns';

	import { earliestConcert } from '$lib/features/concert/concert';
	import type { Event } from '$lib/features/event/event';

	import Button from '$lib/components/ui/Button.svelte';
	import Fader from '$lib/components/Fader.svelte';
	import GlowCursor from '$lib/components/GlowCursor.svelte';

	import TicketIcon from '~icons/mdi/ticket-confirmation-outline';
	import MapIcon from '~icons/mdi/map-marker';
	import CalendarIcon from '~icons/mdi/calendar';
	import MusicIcon from '~icons/mdi/music';

	type Props = {
		event: Event;
		active?: boolean;
		prefix?: string;
	};

	let { event, active, prefix }: Props = $props();
	let fromDate = $derived(earliestConcert(event.concerts)?.from);

	let artists = $derived(event.concerts.map(({ artist }) => artist));

	// https://stackoverflow.com/questions/2218999/how-to-remove-all-duplicates-from-an-array-of-objects
	const genres = $derived(
		artists
			.flatMap((artist) => artist.genres)
			.filter((value, index, self) => index === self.findIndex((t) => t.id === value.id))
			.sort((a, b) => {
				const nameA = a.name.toUpperCase();
				const nameB = b.name.toUpperCase();

				return nameA < nameB ? -1 : nameA > nameB ? 1 : 0;
			})
	);

	let genresString = $derived(genres.map((genre) => genre.name).join(', '));

	const locationUrl = $derived(
		new URL(
			`https://www.google.com/maps/search/?` +
				new URLSearchParams({
					api: '1',
					query: `${event.venue.name},${event.venue.city},${event.venue.countryCode}`
				})
		)
	);
</script>

<div class="relative isolate flex h-[calc((100svh/4)*3)] items-end overflow-hidden pb-16">
	<img
		src={event.imageUrl}
		alt="Event cover"
		class="absolute top-0 left-0 -z-10 h-full w-full object-cover brightness-85"
	/>

	<Fader class="absolute -z-10 h-1/2" direction="up" />

	<div class="px-auto flex w-full flex-col">
		<h1 class="mb-8 text-7xl font-bold">{event.title}</h1>

		<div class="text-text/85 flex">
			<div class="flex flex-1 flex-col gap-1">
				{#if fromDate}
					<div class="flex items-center gap-4">
						<CalendarIcon />
						<span>{format(fromDate, DATE_FORMAT)}</span>
					</div>
				{/if}
				<div class="flex items-center gap-4">
					<MapIcon />
					<a href={locationUrl.toString()}
						>{event.venue.name}, {event.venue.city} ({event.venue.countryCode})</a
					>
				</div>
				<div class="flex items-center gap-4">
					<MusicIcon />
					<span>{genresString}</span>
				</div>
			</div>

			<div class="flex flex-col justify-end gap-2">
				<Button href={event.ticketUrl} class="w-full"><TicketIcon />Køb billet</Button>
				{#if !active}
					<Button href="/events/{event.id}" variant="outline" class="w-full">Læs mere</Button>
				{/if}
			</div>
		</div>
	</div>
</div>
