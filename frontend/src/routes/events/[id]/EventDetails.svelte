<script lang="ts">
	import MapIcon from '~icons/mdi/map-marker';
	import CalendarIcon from '~icons/mdi/calendar';
	import MusicIcon from '~icons/mdi/music';
	import GroupIcon from '~icons/mdi/account-group';
	import { formatDateStr } from '$lib/time';
	import type { Event } from '$lib/event';
	import Button from '$lib/components/Button.svelte';

	type Props = {
		event: Event;
		active?: boolean;
	};

	let { event, active }: Props = $props();
	let earliestEvent = $derived(event.concerts[0]);

	let artistNames = $derived(event.concerts.map((concert) => concert.artist.name));
</script>

<div class="relative h-[calc((100vh/4)*3)] overflow-hidden">
	<img
		src={event?.coverImageUrl}
		alt={event?.title}
		class="absolute top-0 left-0 h-full w-full object-cover"
	/>
	<!-- FADER -->
	<!-- <Fader direction="to-bottom" /> -->
	<div class="px-auto absolute bottom-0 left-0 w-full space-y-2 px-12 pb-12">
		<h1 class="mb-8 w-full text-5xl font-bold md:text-8xl">{event?.title}</h1>
		<div class="flex w-full items-end gap-8">
			<div class="w-full space-y-1 text-zinc-300">
				<div class="flex items-center gap-2">
					<CalendarIcon />
					<span>{formatDateStr(earliestEvent.from)}</span>
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
					<span>{event.venue.name}, {event.venue.city}</span>
				</div>
			</div>
			<div class="w-96 space-y-2">
				<Button expandX expandY extraClass="h-18">Køb billet</Button>
				{#if !active}
					<Button type="secondary" expandX expandY extraClass="h-18">Læs mere</Button>
				{/if}
			</div>
		</div>
	</div>
</div>
