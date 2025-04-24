<script lang="ts">
	import List from '$lib/components/ui/List.svelte';
	import type { Permission } from '$lib/features/auth/permission';
	import type { Event } from '$lib/features/event/event';
	import EventEntry from './EventEntry.svelte';

	type Props = {
		events: Event[];
		memberPermissions: Permission[];
	};

	let { events, memberPermissions }: Props = $props();
</script>

<List>
	{#each events as event (event.id)}
		<EventEntry {event} {memberPermissions} />
	{/each}
</List>
