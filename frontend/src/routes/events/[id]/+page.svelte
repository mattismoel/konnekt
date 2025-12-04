<script lang="ts">
	import { page } from "$app/state";
	import EventCalendar from "$lib/components/EventCalendar.svelte";
	import EventDetails from "$lib/components/EventDetails.svelte";
	import EventGrid from "$lib/components/EventGrid.svelte";
	import { earliestConcert } from "$lib/features/event.js";
	import { getUpcomingEvents } from "$lib/features/event.remote";
	import { DATETIME_FORMAT } from "$lib/time.js";
	import { format } from "date-fns";

	let { data } = $props();
	const fromDate = $derived(earliestConcert(data.event.concerts)?.fromDate);
</script>

<svelte:head>
	<title>Konnekt | Event | {data.event.title}</title>
	<meta
		name="description"
		content="Oplev vores kommende event ${data.event.title} {fromDate
			? format(fromDate, DATETIME_FORMAT)
			: ''}"
	/>
</svelte:head>

<main class="flex min-h-svh flex-col pb-16">
	<EventDetails active={page.params.id === data.event.id} event={data.event} />

	<div class="border-t border-t-zinc-900">
		<article class="mx-responsive w-full space-y-16 pt-16 pb-16">
			<section class="prose max-w-none prose-invert">
				{@html data.event.description}
			</section>

			<EventCalendar event={data.event} />

			{#await getUpcomingEvents({ filter: `id!="${data.event.id}"` }) then upcomingEvents}
				{#if upcomingEvents.items.length > 0}
					<section>
						<h1 class="mb-8 text-2xl font-bold">Se også</h1>
						<EventGrid events={upcomingEvents.items} />
					</section>
				{/if}
			{/await}
		</article>
	</div>
</main>
