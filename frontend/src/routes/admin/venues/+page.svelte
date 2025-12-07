<script lang="ts">
	import AdminPageHeader from "$lib/components/AdminPageHeader.svelte";
	import SearchBar from "$lib/components/SearchBar.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import { hasPermissions } from "$lib/features/auth/auth.remote";
	import VenueListEntry from "$lib/features/venues/components/VenueListEntry.svelte";
	import { getVenues } from "$lib/features/venues/venue.remote";
	import { EntrySearcher } from "$lib/search.svelte";

	const { items: venues } = await getVenues(undefined);

	const searcher = new EntrySearcher(venues, "name");
</script>

<main class="mx-responsive py-32">
	<AdminPageHeader title="Venues" description="Overblik over venues.">
		{#if await hasPermissions(["venues:edit"])}
			<Button href="/admin/venues/create">+ Tilføj</Button>
		{/if}
	</AdminPageHeader>

	<SearchBar bind:search={searcher.search} class="mb-4 w-full" />

	<ul>
		{#each searcher.results as venue}
			<VenueListEntry {venue} />
		{/each}
	</ul>
</main>
