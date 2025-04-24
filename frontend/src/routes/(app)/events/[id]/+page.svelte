<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCalendar from '$lib/components/event-calendar/EventCalendar.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import EventDetails from './EventDetails.svelte';

	const { data } = $props();
	const { event, recommendedEvents } = $derived(data);

	const eventArtists = $derived(event.concerts.map((c) => c.artist));
</script>

<main class="min-h-sub-nav flex flex-col gap-16 pb-16 text-white">
	<EventDetails active {event} />
	<article class="px-auto space-y-16 pt-8 pb-16">
		<section class="prose prose-invert max-w-none">
			{@html event.description}
		</section>
		<EventCalendar {event} />
		<section>
			<h1 class="mb-4 text-2xl font-bold">Se også</h1>
			<Caroussel>
				{#each recommendedEvents as event (event.id)}
					<EventCard {event} />
				{/each}
			</Caroussel>
		</section>
	</article>
</main>
