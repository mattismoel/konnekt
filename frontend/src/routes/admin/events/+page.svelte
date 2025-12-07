<script lang="ts">
	import AdminPageHeader from "$lib/components/AdminPageHeader.svelte";
	import SearchBar from "$lib/components/SearchBar.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import type { Artist } from "$lib/features/artists/artist";
	import { type Event } from "$lib/features/events/event";
	import { getPreviousEvents, getUpcomingEvents } from "$lib/features/events/event.remote";
	import { EntrySearcher } from "$lib/search.svelte";
	import { DATETIME_FORMAT } from "$lib/time.js";
	import { format } from "date-fns";

	const { items: upcomingEvents } = await getUpcomingEvents(undefined);
	const { items: previousEvents } = await getPreviousEvents(undefined);

	let searcher = new EntrySearcher(upcomingEvents, "title");

	const formatArtists = (artists: Artist[]): string => {
		if (artists.length > 2) {
			return (
				artists
					.slice(0, 2)
					.map((a) => a.name)
					.join(", ") +
				" " +
				`(+${artists.length - 2} mere)`
			);
		}

		return artists.map((a) => a.name).join(", ");
	};
</script>

{#snippet entry(event: Event)}
	{@const artists = event.concerts.map((concert) => concert.artist)}

	<li
		class="group flex items-center rounded-xl border border-zinc-800 bg-zinc-900 transition-colors hover:border-zinc-700 hover:bg-zinc-800"
	>
		<a href="/admin/events/edit/{event.id}" class="w-full py-4 pl-8">
			<div>
				<p class="font-medium transition-colors group-hover:text-text-light group-hover:underline">
					{event.title}
				</p>
				<div class="text-base">
					<p>{format(event.concerts[0].fromDate, DATETIME_FORMAT)}</p>
					<p>{formatArtists(artists)}</p>
				</div>
			</div>
		</a>
	</li>
{/snippet}

<main class="mx-responsive min-h-svh py-32">
	<AdminPageHeader title="Events" description="Overblik over alle events.">
		<Button href="/admin/events/create">Tilføj</Button>
	</AdminPageHeader>

	<div>
		<SearchBar bind:search={searcher.search} class="mb-8 w-full" />
		{#if searcher.search}
			<ul>
				{#each searcher.results as event}
					{@render entry(event)}
				{/each}
			</ul>
		{:else}
			<div class="mb-8">
				<h3 class="mb-2 font-medium">Kommende events</h3>
				{#if upcomingEvents.length > 0}
					<ul>
						{#each upcomingEvents as event}
							{@render entry(event)}
						{/each}
					</ul>
				{:else}
					<p>Ingen kommende events...</p>
				{/if}
			</div>

			{#if previousEvents.length > 0}
				<div>
					<h3 class="mb-2 font-medium">Tidligere events</h3>
					<ul>
						{#each previousEvents as event}
							{@render entry(event)}
						{/each}
					</ul>
				</div>
			{/if}
		{/if}
	</div>
</main>
