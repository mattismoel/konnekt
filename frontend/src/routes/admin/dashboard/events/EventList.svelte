<script lang="ts">
	import type { Permission } from '$lib/features/auth/permission';
	import type { Event } from '$lib/features/event/event';
	import EventEntry from './EventEntry.svelte';

	type Props = {
		events: Event[];
		userPermissions: Permission[];
		onDelete: (id: number) => void;
	};

	let { events, userPermissions, onDelete }: Props = $props();
</script>

<ul class="space-y-2">
	{#each events as event (event.id)}
		<EventEntry {event} {userPermissions} onDelete={() => onDelete(event.id)} />
	{/each}
</ul>
