<script lang="ts">
	import MapIcon from '~icons/mdi/map-marker';
	import CalendarIcon from '~icons/mdi/calendar';
	import MusicIcon from '~icons/mdi/music';
	import GroupIcon from '~icons/mdi/account-group';
	import Button from '$lib/components/ui/Button.svelte';
	import { formatDateStr } from '$lib/time';
	import type { Event } from '$lib/event';
	import { earliestConcert } from '$lib/concert';
	import Fader from '$lib/components/ui/Fader.svelte';

	type Props = {
		event: Event;
		active?: boolean;
		prefix?: string;
	};

	let { event, active, prefix }: Props = $props();
	let fromDate = $derived(earliestConcert(event.concerts)?.from);

	let artistNames = $derived(event.concerts.map((concert) => concert.artist.name));
</script>

<div class="relative h-[calc((100vh/4)*3)] overflow-hidden">
	<img
		src={event?.imageUrl}
		alt={event?.title}
		class="absolute top-0 left-0 h-full w-full object-cover"
	/>
	<!-- FADER -->
	<Fader direction="up" class="absolute h-[512px] from-zinc-950 md:h-96" />
	<div class="px-auto absolute bottom-0 left-0 flex w-full flex-col gap-y-2 px-12 pb-12">
		<span class="mb-1 font-medium">{prefix}</span>
		<h1 class="mb-4 w-full text-5xl font-bold md:text-8xl">{event?.title}</h1>
		<div class="flex w-full flex-col items-end gap-8 md:flex-row">
			<div class="w-full space-y-1 text-zinc-300">
				<div class="flex items-center gap-2">
					<CalendarIcon />
					<time>{formatDateStr(fromDate || new Date())}</time>
				</div>
				<div class="flex items-center gap-2">
					<GroupIcon />
					<span>{artistNames.join(', ')}</span>
				</div>
				<div class="flex items-center gap-2">
					<MusicIcon />
					<span>{['EDM', 'Dance', 'House', "Drum'n'Bass"].join(', ')}</span>
				</div>
				<div class="flex items-center gap-2">
					<MapIcon />
					<address class="not-italic">{event.venue.name}, {event.venue.city}</address>
				</div>
			</div>
			<div class="w-full space-y-2 md:w-96">
				<form action={event.ticketUrl}>
					<Button type="submit" expandX expandY class="h-18">Køb billet</Button>
				</form>
				{#if !active}
					<form action="/events/{event.id}">
						<Button type="submit" variant="secondary" expandX expandY class="h-18">Læs mere</Button>
					</form>
				{/if}
			</div>
		</div>
	</div>
</div>
