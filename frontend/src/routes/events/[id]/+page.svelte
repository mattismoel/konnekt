<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import EventDetails from './EventDetails.svelte';

	const { data } = $props();
	const { event, recommendedEvents } = $derived(data);
</script>

<main class="min-h-sub-nav pb-16 text-white">
	<EventDetails active {event} />
	<article class="px-auto prose prose-invert max-w-none pt-20 pb-16 text-gray-400">
		{@html event.description}
	</article>
	<div class="px-auto">
		<Caroussel>
			{#each recommendedEvents as event (event.id)}
				<EventCard {event} />
			{/each}
		</Caroussel>
	</div>
</main>
