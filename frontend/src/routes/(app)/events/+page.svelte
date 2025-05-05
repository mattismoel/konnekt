<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import EventDetails from './[id]/EventDetails.svelte';

	let { data } = $props();
	let { events } = $derived(data);
</script>

<main class="min-h-svh">
	{#if events.length > 0}
		<EventDetails event={events[0]} prefix="Næste event:" />
	{/if}
	<div class="px-auto flex flex-col gap-16 pt-32 pb-16">
		<section class="flex flex-col">
			<h1 class="font-heading mb-4 text-5xl font-bold md:text-7xl">Events</h1>
			<span class="text-text/75"> Her kan du se alle kommende events. </span>
		</section>
		{#if events.length <= 0}
			<span class="text-text/50 italic">Der er ingen kommende events i øjeblikket...</span>
		{/if}
		<Caroussel>
			{#each events as event (event.id)}
				<EventCard {event} />
			{/each}
		</Caroussel>
	</div>
</main>
