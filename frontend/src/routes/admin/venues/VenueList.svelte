<script lang="ts">
	import SearchList from '$lib/components/SearchList.svelte';
	import type { Permission } from '$lib/features/auth/permission';
	import type { Venue } from '$lib/features/venue/venue';
	import VenueEntry from './VenueEntry.svelte';

	type Props = {
		venues: Venue[];
	};

	let { venues }: Props = $props();

	let search = $state('');

	let filteredVenues = $derived(
		venues.filter((v) => v.name.toLowerCase().includes(search.toLowerCase()))
	);
</script>

<SearchList bind:search>
	{#each filteredVenues as venue (venue.id)}
		<VenueEntry {venue} />
	{/each}
</SearchList>
