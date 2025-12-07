<script lang="ts">
	import EventDetails from "$lib/features/events/components/EventDetails.svelte";
	import EventGrid from "$lib/features/events/components/EventGrid.svelte";
	import { getUpcomingEvents } from "$lib/features/events/event.remote.js";

	const { items: events } = await getUpcomingEvents({ filter: "isPublic=true" });
	const eventNames = $derived(events.map((e) => e.title).join(", "));
</script>

<svelte:head>
	<title>Konnekt | Events</title>
	<meta
		name="description"
		content="Her kan du se Konnekts kommende events. Oplev blandt andet events som ${eventNames}"
	/>
</svelte:head>

{#if events.length > 0}
	<main class="min-h-svh">
		<EventDetails event={events[0]} prefix="Næste event:" />

		<div class="border-t border-t-zinc-900">
			<div class="mx-responsive flex flex-col pt-16 pb-16 md:pt-16">
				<h1 class="font-heading mb-8 text-4xl font-bold text-text-light">Kommende events</h1>
				<EventGrid {events} />
			</div>
		</div>
	</main>
{:else}
	<main class="mx-responsive flex min-h-svh flex-col items-center justify-center">
		<span class="text-text/75 text-center italic">
			Der er ingen aktuelle events i øjeblikket...
		</span>
	</main>
{/if}
