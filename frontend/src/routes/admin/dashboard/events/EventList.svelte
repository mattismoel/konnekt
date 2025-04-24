<script lang="ts">
	import List from '$lib/components/ui/List.svelte';
	import type { Permission } from '$lib/features/auth/permission';
	import type { Event } from '$lib/features/event/event';
	import EventEntry from './EventEntry.svelte';

	type Props = {
		events: Event[];
		memberPermissions: Permission[];
		onDelete: (id: number) => void;
	};

	let { events, memberPermissions, onDelete }: Props = $props();
</script>

<List>
	{#each events as event (event.id)}
		<EventEntry {event} {memberPermissions} onDelete={() => onDelete(event.id)} />
	{/each}
</List>
