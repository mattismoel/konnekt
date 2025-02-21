<script lang="ts">
	import type { Event } from '$lib/event';
	import EventCard from './EventCard.svelte';

	type Props = {
		title?: string;
		events: Event[];
	};

	let { events, title = 'Events.' }: Props = $props();
</script>

<h3 class="mb-4 text-2xl font-bold">{title}</h3>
<div class="flex gap-2">
	{#each events as event (event.id)}
		<EventCard {event} />
	{/each}
</div>
