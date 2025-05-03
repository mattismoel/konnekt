<script lang="ts">
	import SearchBar from '$lib/components/SearchBar.svelte';
	import List from '$lib/components/ui/List.svelte';
	import type { Permission } from '$lib/features/auth/permission';
	import type { Venue } from '$lib/features/venue/venue';
	import VenueEntry from './VenueEntry.svelte';

	type Props = {
		venues: Venue[];
		memberPermissions: Permission[];
	};

	let { venues, memberPermissions }: Props = $props();

	let search = $state('');

	let filteredVenues = $derived(
		venues.filter((v) => v.name.toLowerCase().includes(search.toLowerCase()))
	);
</script>

<div class="flex flex-col gap-8">
	<SearchBar bind:value={search} />
	<List>
		{#each filteredVenues as venue (venue.id)}
			<VenueEntry {venue} {memberPermissions} />
		{/each}
	</List>
</div>
