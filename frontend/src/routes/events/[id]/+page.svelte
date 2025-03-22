<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCalendar from '$lib/components/EventCalendar.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import EventDetails from './EventDetails.svelte';

	const { data } = $props();
	const { event, recommendedEvents } = $derived(data);

	const eventArtists = $derived(event.concerts.map((c) => c.artist));
</script>

<main class="min-h-sub-nav pb-16 text-white">
	<EventDetails active {event} />
	<section>
		<article class="px-auto prose prose-invert max-w-none pt-8 pb-16 text-gray-400">
			{@html event.description}
		</article>
		<EventCalendar {event} />
		<ul>
			{#each eventArtists as artist (artist.id)}
				<li><a href="/artists/{artist.id}">{artist.name}</a></li>
			{/each}
		</ul>
	</section>
	<section class="px-auto">
		<h1 class="mb-4 text-2xl font-bold">Se også</h1>
		<Caroussel>
			{#each recommendedEvents as event (event.id)}
				<EventCard {event} />
			{/each}
		</Caroussel>
	</section>
</main>
