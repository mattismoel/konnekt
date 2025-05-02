<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import VenueEntry from './VenueEntry.svelte';
	import SearchBar from '$lib/components/SearchBar.svelte';
	import { hasPermissions } from '$lib/features/auth/permission';
	import DashboardLayout from '../DashboardLayout.svelte';
	import DashboardHeader from '../DashboardHeader.svelte';
	import HeaderActions from '../HeaderActions.svelte';

	let { data } = $props();

	let search = $state('');

	let venues = $derived(
		data.venues.filter((v) => v.name.toLowerCase().includes(search.toLowerCase()))
	);
</script>

<DashboardLayout>
	<DashboardHeader
		title="Venues"
		description="Overblik over alle venues, som er associerede med events for Konnekt."
	>
		<HeaderActions>
			<Button
				href="/admin/venues/create"
				disabled={!hasPermissions(data.member.permissions, ['edit:venue'])}
			>
				<PlusIcon />Tilføj
			</Button>
		</HeaderActions>
	</DashboardHeader>

	{#if hasPermissions(data.member.permissions, ['view:venue'])}
		<section class="space-y-4">
			<SearchBar bind:value={search} />
			<ul>
				{#each venues as venue (venue.id)}
					<VenueEntry {venue} memberPermissions={data.member.permissions} />
				{/each}
			</ul>
		</section>
	{:else}
		<span>Du har ikke tilladelse til at se venues...</span>
	{/if}
</DashboardLayout>
