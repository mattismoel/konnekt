<script lang="ts">
	import PlusIcon from '~icons/mdi/plus';
	import EventEntry from './EventEntry.svelte';

	let { data } = $props();

	let { upcomingEvents } = $derived(data);
</script>

<div class="rounded-sm border border-zinc-800 bg-zinc-900 p-8">
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
</div>
