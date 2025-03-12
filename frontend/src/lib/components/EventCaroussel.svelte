<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { Event } from '$lib/event';
	import EventCard from './EventCard.svelte';
	import Fader from './ui/Fader.svelte';

	type Props = {
		title?: string;
		events: Event[];
	};

	let { events }: Props = $props();

	let scrollX = $state(0);
	let inner: HTMLDivElement;
	let outer: HTMLDivElement;

	let isScrolledToRight = $state(false);

	const updateScroll = (newScrollX: number) => {
		const innerWidth = inner.scrollWidth;
		const outerWidth = outer.getBoundingClientRect().width;

		const diff = innerWidth - outerWidth;
		scrollX = newScrollX;

		isScrolledToRight = scrollX >= diff;
	};
</script>

<div class="relative isolate" bind:this={outer}>
	<!--RIGHT FADER -->
	<Fader
		direction="left"
		class={cn('absolute z-10 w-32 from-zinc-950 transition-colors duration-300', {
			'from-transparent': isScrolledToRight
		})}
	/>
	<Fader
		direction="right"
		class={cn('absolute z-10 w-32 from-zinc-950 transition-colors duration-300', {
			'from-transparent': scrollX <= 0
		})}
	/>
	<div
		bind:this={inner}
		class="flex w-full gap-4 overflow-x-scroll"
		onscroll={(e) => updateScroll(e.currentTarget.scrollLeft)}
	>
		{#each events as event (event.id)}
			<EventCard {event} />
		{/each}
	</div>
</div>
