<script lang="ts">
	import AdminPageHeader from "$lib/components/AdminPageHeader.svelte";
	import SearchBar from "$lib/components/SearchBar.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import EventListEntry from "$lib/features/events/components/EventListEntry.svelte";
	import { getPreviousEvents, getUpcomingEvents } from "$lib/features/events/event.remote";
	import { EntrySearcher } from "$lib/search.svelte";

	const { items: upcomingEvents } = await getUpcomingEvents(undefined);
	const { items: previousEvents } = await getPreviousEvents(undefined);

	let searcher = new EntrySearcher(upcomingEvents, "title");
</script>

<main class="mx-responsive min-h-svh py-32">
	<AdminPageHeader title="Events" description="Overblik over alle events.">
		<Button href="/admin/events/create">Tilføj</Button>
	</AdminPageHeader>

	<div>
		<SearchBar bind:search={searcher.search} class="mb-8 w-full" />
		{#if searcher.search}
			<ul>
				{#each searcher.results as event}
					<EventListEntry {event} />
				{/each}
			</ul>
		{:else}
			<div class="mb-16">
				<h3 class="mb-4 text-2xl font-bold">Kommende events</h3>
				{#if upcomingEvents.length > 0}
					<ul>
						{#each upcomingEvents as event}
							<EventListEntry {event} />
						{/each}
					</ul>
				{:else}
					<p class="italic">Ingen kommende events...</p>
				{/if}
			</div>

			{#if previousEvents.length > 0}
				<div>
					<h3 class="mb-4 text-2xl font-bold">Tidligere events</h3>
					<ul>
						{#each previousEvents as event}
							<EventListEntry {event} />
						{/each}
					</ul>
				</div>
			{/if}
		{/if}
	</div>
</main>
