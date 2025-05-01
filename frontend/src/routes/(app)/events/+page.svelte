<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import EventDetails from './[id]/EventDetails.svelte';

	let { data } = $props();
	let { events } = $derived(data);
</script>

<main class="min-h-sub-nav">
	{#if events.length > 0}
		<EventDetails event={events[0]} prefix="Næste event:" />
	{/if}
	<div class="px-auto py-16">
		<h1 class="mb-4 text-3xl font-bold">Events</h1>
		<Caroussel>
			{#each events as event (event.id)}
				<EventCard {event} />
			{/each}
		</Caroussel>
	</div>
</main>
