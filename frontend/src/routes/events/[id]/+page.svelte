<script lang="ts">
	import { page } from "$app/state";
	import EventCalendar from "$lib/features/events/components/EventCalendar.svelte";
	import EventDetails from "$lib/features/events/components/EventDetails.svelte";
	import EventGrid from "$lib/features/events/components/EventGrid.svelte";
	import { earliestConcert } from "$lib/features/events/event";
	import { getEvent, getUpcomingEvents } from "$lib/features/events/event.remote";
	import { DATETIME_FORMAT } from "$lib/time.js";
	import { format } from "date-fns";

	const event = await getEvent(page.params.id!);
	let { id, title, description, concerts } = event;

	const fromDate = $derived(earliestConcert(concerts)?.fromDate);

	const { items: upcomingEvents } = await getUpcomingEvents({
		filter: `isPublic=true && id!="${id}"`
	});
</script>

<svelte:head>
	<title>Konnekt | Event | {title}</title>
	<meta
		name="description"
		content="Oplev vores kommende event ${title} {fromDate
			? format(fromDate, DATETIME_FORMAT)
			: ''}"
	/>
</svelte:head>

<main class="flex min-h-svh flex-col pb-16">
	<EventDetails active={page.params.id === id} {event} />

	<div class="border-t border-t-zinc-900">
		<article class="mx-responsive w-full space-y-16 pt-16 pb-16">
			<section class="prose max-w-none prose-invert">
				{@html description}
			</section>

			<EventCalendar {event} />

			{#if upcomingEvents.length > 0}
				<section>
					<h1 class="mb-8 text-2xl font-bold">Se også</h1>
					<EventGrid events={upcomingEvents} />
				</section>
			{/if}
		</article>
	</div>
</main>
