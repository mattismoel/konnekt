<script lang="ts">
	import MapIcon from '~icons/ph/map-pin';
	import CalendarIcon from '~icons/ph/calendar';
	import MusicIcon from '~icons/ph/music-notes';

	import type { Event } from '$lib/event';

	import { cn } from '$lib/clsx';
	import { formatDateStr } from '$lib/time';

	type Props = {
		event: Event;
	};

	const { event }: Props = $props();

	const earliestEvent = $derived(event.concerts[0]);
</script>

<div class="relative h-[calc((100vh/4)*3)] overflow-hidden">
	<img
		src={event?.coverImageUrl}
		alt={event?.title}
		class="absolute top-0 left-0 h-full w-full object-cover"
	/>
	<!-- FADER -->
	<!-- <Fader direction="to-bottom" /> -->
	<div class="absolute bottom-0 left-0 space-y-2 px-12 pb-12">
		<h1 class={cn('mb-2 w-full text-5xl font-bold md:text-8xl')}>{event?.title}</h1>
		<div class="flex items-center gap-2">
			<MapIcon />
			<span>{event.venue.name}, {event.venue.city}</span>
		</div>
		<div class="flex items-center gap-2">
			<CalendarIcon />
			<span>{formatDateStr(earliestEvent.from)}</span>
		</div>
		<div class="flex items-center gap-2">
			<MusicIcon />
			<!-- <span>{event?.genres.join(', ')}</span> -->
		</div>
	</div>
</div>
