<script lang="ts">
	import PlusIcon from '~icons/mdi/plus';
	import EventEntry from './EventEntry.svelte';
	import Card from '$lib/components/ui/Card.svelte';

	let { data } = $props();

	let { upcomingEvents } = $derived(data);

	$inspect(upcomingEvents)
</script>

<Card>
	<div class="flex justify-between">
		<h1 class="mb-4 text-2xl font-bold">Kommende events.</h1>
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
