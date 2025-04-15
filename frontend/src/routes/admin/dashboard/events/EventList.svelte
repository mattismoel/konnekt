<script lang="ts">
	import type { Event } from '$lib/features/event/event';
	import EventEntry from './EventEntry.svelte';

	type Props = {
		events: Event[];
		onDelete: (id: number) => void;
	};

	let { events, onDelete }: Props = $props();
</script>

<ul class="space-y-2">
	{#each events as event (event.id)}
		<EventEntry {event} onDelete={() => onDelete(event.id)} />
	{/each}
</ul>
