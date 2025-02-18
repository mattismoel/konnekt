<script lang="ts">
	import EventDetails from './EventDetails.svelte';

	const { data } = $props();
	const { event } = $derived(data);
</script>

<main class="min-h-sub-nav bg-black pb-16 text-white">
	<EventDetails {event} />
	<article class="px-auto prose prose-invert max-w-none pt-20 pb-16 text-gray-400">
		{@html event.description}
	</article>
	<div class="px-auto">
		<h2 class="mb-6 text-2xl font-bold">Billetten gælder også til.</h2>
		<!-- <EventCalendar
			activeID={event?.id}
			events={weeklyEventsResult?.events || []}
			class="h-64 mb-12"
		/> -->
		<h1 class="mb-12 text-4xl font-bold">Se også.</h1>
		<!-- <EventCaroussel events={recommendedEventsResult?.events || []} /> -->
	</div>
</main>
