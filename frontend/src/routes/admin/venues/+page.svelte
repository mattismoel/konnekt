<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import { hasPermissions } from '$lib/features/auth/permission';
	import DashboardLayout from '../DashboardLayout.svelte';
	import DashboardHeader from '../DashboardHeader.svelte';
	import HeaderActions from '../HeaderActions.svelte';
	import VenueList from './VenueList.svelte';

	let { data } = $props();
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
		<VenueList venues={data.venues} memberPermissions={data.member.permissions} />
	{:else}
		<span>Du har ikke tilladelse til at se venues...</span>
	{/if}
</DashboardLayout>
