<script lang="ts">
	import SearchBar from '$lib/components/SearchBar.svelte';
	import List from '$lib/components/ui/List.svelte';
	import type { Permission } from '$lib/features/auth/permission';
	import type { Event } from '$lib/features/event/event';
	import EventEntry from './EventEntry.svelte';

	type Props = {
		previousEvents: Event[];
		upcomingEvents: Event[];
		memberPermissions: Permission[];
	};

	let { previousEvents, upcomingEvents, memberPermissions }: Props = $props();

	let search = $state('');

	let filteredEvents = $derived(
		[...previousEvents, ...upcomingEvents].filter((event) =>
			event.title.toLowerCase().includes(search.toLowerCase())
		)
	);
</script>

<div class="space-y-8">
	<SearchBar bind:value={search} />

	{#if search}
		<List>
			{#each filteredEvents as event (event.id)}
				<EventEntry {event} {memberPermissions} />
			{/each}
		</List>
	{:else}
		<List>
			{#each upcomingEvents as event (event.id)}
				<EventEntry {event} {memberPermissions} />
			{/each}
		</List>
		<details>
			<summary class="mb-4">Tidligere events ({previousEvents.length})</summary>
			{#each previousEvents as event (event.id)}
				<EventEntry {event} {memberPermissions} />
			{/each}
		</details>
	{/if}
</div>
