<script lang="ts">
	import EventCaroussel from '$lib/components/EventCaroussel.svelte';
	import EventDetails from './EventDetails.svelte';

	const { data } = $props();
	const { event, recommendedEvents } = $derived(data);
</script>

<main class="min-h-sub-nav bg-black pb-16 text-white">
	<EventDetails active {event} />
	<article class="px-auto prose prose-invert max-w-none pt-20 pb-16 text-gray-400">
		{@html event.description}
	</article>
	<div class="px-auto">
		<EventCaroussel title="Se også." events={recommendedEvents} />
	</div>
</main>
