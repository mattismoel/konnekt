<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCalendar from '$lib/components/event-calendar/EventCalendar.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import EventDetails from './EventDetails.svelte';

	const { data } = $props();
</script>

<main class="min-h-sub-nav flex flex-col gap-16 pb-16 text-white">
	<EventDetails active event={data.event} />
	<article class="px-auto space-y-16 pt-8 pb-16">
		<section class="prose prose-invert max-w-none">
			{@html data.event.description}
		</section>
		<EventCalendar event={data.event} />
		{#if data.recommendedEvents.length > 0}
			<section>
				<h1 class="mb-4 text-2xl font-bold">Se også</h1>
				<Caroussel>
					{#each data.recommendedEvents as event (event.id)}
						<EventCard {event} />
					{/each}
				</Caroussel>
			</section>
		{/if}
	</article>
</main>
