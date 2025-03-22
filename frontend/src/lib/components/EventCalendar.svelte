<script lang="ts">
	import type { Event } from '$lib/event';
	import { differenceInMinutes } from 'date-fns';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		event: Event;
	};

	let { event, ...rest }: Props = $props();

	let concerts = $derived(event.concerts.sort((a, b) => a.from.getTime() - b.from.getTime()));

	let totalDurationMinutes = $derived.by(() => {
		if (concerts.length <= 0) return 0;

		return differenceInMinutes(concerts[concerts.length - 1].to, concerts[0].from);
	});

	$inspect(totalDurationMinutes, concerts);
</script>

<div
	class="relative h-64 w-full max-w-xl border border-zinc-900 bg-radial from-zinc-900 to-zinc-950 p-4"
>
	{#each concerts as concert (concert.id)}
		{@const concertDurationMinutes = differenceInMinutes(concert.to, concert.from)}
		<div
			style:height="calc(({concertDurationMinutes}/{totalDurationMinutes})*100%)"
			style:top="calc(({differenceInMinutes(concert.to, concert.from)} / {totalDurationMinutes})*100%)"
			class="absolute w-full bg-red-200"
		></div>
	{/each}
</div>
