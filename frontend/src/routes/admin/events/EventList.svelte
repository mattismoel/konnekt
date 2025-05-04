<script lang="ts">
	import SearchList from '$lib/components/SearchList.svelte';
	import type { Event } from '$lib/features/event/event';
	import EventEntry from './EventEntry.svelte';

	type Props = {
		previousEvents: Event[];
		upcomingEvents: Event[];
	};

	let { previousEvents, upcomingEvents }: Props = $props();

	let search = $state('');

	let filteredEvents = $derived(
		[...previousEvents, ...upcomingEvents].filter((event) =>
			event.title.toLowerCase().includes(search.toLowerCase())
		)
	);
</script>

<SearchList bind:search>
	{#if search}
		{#each filteredEvents as event (event.id)}
			<EventEntry {event} />
		{/each}
	{:else}
		{#each upcomingEvents as event (event.id)}
			<EventEntry {event} />
		{/each}
		<details>
			<summary class="mb-4">Tidligere events ({previousEvents.length})</summary>
			{#each previousEvents as event (event.id)}
				<EventEntry {event} />
			{/each}
		</details>
	{/if}
</SearchList>
