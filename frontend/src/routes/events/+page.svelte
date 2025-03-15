<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCard from '$lib/components/EventCard.svelte';

	let { data } = $props();
	let { events } = $derived(data);
</script>

<main class="min-h-sub-nav pt-20">
	<div class="px-auto">
		<h1 class="mb-4 text-3xl font-bold">Events.</h1>
		<Caroussel>
			{#each events as event (event.id)}
				<EventCard {event} />
			{/each}
		</Caroussel>
	</div>
</main>
