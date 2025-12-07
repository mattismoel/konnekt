<script lang="ts">
	import type { Event } from "$lib/features/events/event";
	import EventEntry from "./EventEntry.svelte";

	type Props = {
		events: Event[];
	};

	let { events }: Props = $props();
</script>

<div class="@container">
	<div class="grid grid-cols-1 gap-4 @2xl:grid-cols-2 @5xl:grid-cols-3 @7xl:grid-cols-4">
		{#each events as event}
			<EventEntry {event} />
		{/each}
	</div>
</div>
