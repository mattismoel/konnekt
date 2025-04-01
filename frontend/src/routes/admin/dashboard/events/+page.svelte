<script lang="ts">
	import Card from '$lib/components/ui/Card.svelte';
	import EventEntry from './EventEntry.svelte';
	import PlusIcon from '~icons/mdi/plus';

	let { data } = $props();

	let { upcomingEvents } = $derived(data);
</script>

<Card>
	<div class="flex justify-between">
		<h1 class="font-heading mb-4 text-4xl font-bold">Kommende events</h1>
		<a href="/admin/events/edit" class="flex gap-2">
			<PlusIcon class="text-xl" />
			Tilføj event
		</a>
	</div>
	<div class="divide-y">
		{#each upcomingEvents as event (event.id)}
			<EventEntry {event} />
		{/each}
	</div>
</Card>
