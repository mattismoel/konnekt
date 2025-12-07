<script lang="ts">
	import { getUpcomingEvents } from "$lib/features/events/event.remote";
	import Button from "$lib/components/ui/Button.svelte";
	import AdminPageHeader from "$lib/components/AdminPageHeader.svelte";
	import EventListEntry from "$lib/features/events/components/EventListEntry.svelte";

	const { items: upcomingEvents } = await getUpcomingEvents(undefined);
</script>

<main class="mx-responsive py-32">
	<AdminPageHeader
		title="Kommende events"
		description="Her har du overblik over alle aktuelle events, kunstnere og venues."
	/>

	<div class="flex flex-col gap-16">
		<div>
			<div class="mb-8 flex gap-2">
				<Button href="/admin/events/create">+ Tilføj</Button>
				<Button href="/admin/events" variant="secondary">Vis alle</Button>
			</div>
			{#if upcomingEvents.length > 0}
				<ul class="flex flex-col gap-2">
					{#each upcomingEvents as event}
						<EventListEntry {event} />
					{/each}
				</ul>
			{:else}
				<p class="italic">Der er ingen kommende events...</p>
			{/if}
		</div>
	</div>
</main>
