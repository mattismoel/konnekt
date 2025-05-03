<script lang="ts">
	import { differenceInMinutes, startOfHour, addHours, format } from 'date-fns';

	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	import type { Event } from '$lib/features/event/event';

	import ConcertEntry from './ConcertEntry.svelte';

	type Props = HTMLAttributes<HTMLDivElement> & {
		event: Event;
	};

	let { event, ...rest }: Props = $props();

	let concerts = $derived(event.concerts.sort((a, b) => a.from.getTime() - b.from.getTime()));

	let startHour = $derived(
		concerts.length > 0 ? startOfHour(concerts[0].from) : startOfHour(new Date())
	);

	let endHour = $derived(
		concerts.length > 0
			? addHours(startOfHour(concerts[concerts.length - 1].to), 1)
			: addHours(startOfHour(new Date()), 4)
	);

	let timeMarkers = $derived.by(() => {
		let markers: Date[] = [];
		let current = startHour;

		while (current <= endHour) {
			markers = [...markers, current];
			current = addHours(current, 1);
		}

		return markers;
	});

	let totalMinutes = $derived(Math.max(1, differenceInMinutes(endHour, startHour)));
</script>

<div class="flex w-full flex-col gap-8" {...rest}>
	<div>
		<h3 class="mb-2 text-xl font-bold">Program for {event.title}</h3>
		<span class="text-zinc-300">{format(concerts[0].from, 'EE, dd/MM/yyyy')}</span>
	</div>

	<div class="overflow-y-scroll">
		<div class="grid h-full min-h-72 flex-1 grid-cols-[48px_1fr] gap-4">
			<!-- Timeline -->
			<div class="relative">
				{#each timeMarkers as marker, i}
					{@const markerOffset = differenceInMinutes(marker, startHour)}
					<div
						style:top="calc({(markerOffset / totalMinutes) * 100}%)"
						class={cn('text-text/50 absolute flex w-full flex-col gap-1 text-sm', {
							hidden: i === timeMarkers.length - 1
						})}
					>
						<div class="h-[1px] w-full bg-zinc-800"></div>
						<span>{format(marker, 'HH:mm')}</span>
					</div>
				{/each}
			</div>

			<!-- Events -->
			<div class="relative">
				{#each concerts as concert (concert.id)}
					<ConcertEntry {concert} {totalMinutes} {startHour} />
				{/each}
			</div>
		</div>
	</div>
</div>
