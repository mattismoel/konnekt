<script lang="ts">
	import List from '$lib/components/ui/List.svelte';
	import type { Permission } from '$lib/features/auth/permission';
	import type { Venue } from '$lib/features/venue/venue';
	import VenueEntry from './VenueEntry.svelte';

	type Props = {
		venues: Venue[];
		memberPermissions: Permission[];
	};

	let { venues, memberPermissions }: Props = $props();
</script>

<List>
	{#each venues as venue (venue.id)}
		<VenueEntry {venue} {memberPermissions} />
	{/each}
</List>
