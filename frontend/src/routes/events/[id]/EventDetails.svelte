<script lang="ts">
	import { earliestConcert } from '$lib/concert';
	import type { Event } from '$lib/event';

	import Button from '$lib/components/ui/Button.svelte';
	import Fader from '$lib/components/ui/Fader.svelte';

	import MapIcon from '~icons/mdi/map-marker';
	import CalendarIcon from '~icons/mdi/calendar';
	import MusicIcon from '~icons/mdi/music';
	import GroupIcon from '~icons/mdi/account-group';
	import { format } from 'date-fns';
	import { DATE_FORMAT } from '$lib/time';
	import GlowCursor from '$lib/components/GlowCursor.svelte';

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
	);
</script>

<section class="relative h-[calc((100vh/4)*3)] overflow-hidden">
	<img
		src={event?.imageUrl}
		alt={event?.title}
		class="absolute top-0 left-0 h-full w-full object-cover brightness-75"
	/>
	<Fader direction="up" class="absolute z-10 h-[512px] from-zinc-950 md:h-96" />
	<GlowCursor class="z-0" />
	<div class="px-auto absolute bottom-0 left-0 z-20 flex w-full flex-col gap-y-2 px-12 pb-12">
		<span class="mb-1 font-medium">{prefix}</span>
		<h1 class="w-full text-5xl font-bold md:text-9xl">{event?.title}</h1>
		<div class="flex w-full flex-col items-end gap-8 md:flex-row">
			<!-- DETAILS -->
			<section class="w-full space-y-1 text-zinc-300">
				<div class="flex items-center gap-2">
					<CalendarIcon />
					<time>{format(fromDate || new Date(), DATE_FORMAT)}</time>
				</div>
				<div class="flex items-center gap-2">
					<MapIcon />
					<address class="not-italic">{event.venue.name}, {event.venue.city}</address>
				</div>
				<div class="flex items-center gap-2">
					<MusicIcon />
					<span>{genres.map(({ name }) => name).join(', ')}</span>
				</div>
			</section>
			<!-- CTA -->
			<section class="w-full space-y-2 md:w-96">
				<form action={event.ticketUrl}>
					<Button type="submit" class="h-18 w-full">Køb billet</Button>
				</form>
				{#if !active}
					<form action="/events/{event.id}">
						<Button type="submit" variant="secondary" class="h-18 w-full">Læs mere</Button>
					</form>
				{/if}
			</section>
		</div>
	</div>
</section>
