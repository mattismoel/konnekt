<script lang="ts">
	import EventCard from '$lib/components/EventCard.svelte';
	import EventCaroussel from '$lib/components/EventCaroussel.svelte';

	let { data } = $props();
	let { events } = $derived(data);
</script>

<main class="min-h-sub-nav pt-20">
	<div class="px-auto">
		<h1 class="mb-4 text-3xl font-bold">Events.</h1>
		<EventCaroussel {events} />
	</div>
</main>
