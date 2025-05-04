<script lang="ts">
	import PlusIcon from '~icons/mdi/plus';
	import Button from '$lib/components/ui/Button.svelte';

	import { hasPermissions } from '$lib/features/auth/permission';
	import EventList from './EventList.svelte';
	import DashboardLayout from '../DashboardLayout.svelte';
	import DashboardHeader from '../DashboardHeader.svelte';
	import HeaderActions from '../HeaderActions.svelte';
	import { authStore } from '$lib/auth.svelte';

	let { data } = $props();
</script>

<DashboardLayout>
	<DashboardHeader title="Events" description="Overblik over alle events.">
		<HeaderActions>
			<Button
				href="/admin/events/create"
				disabled={!hasPermissions(authStore.permissions, ['edit:event'])}
			>
				<PlusIcon />Tilføj
			</Button>
		</HeaderActions>
	</DashboardHeader>

	{#if hasPermissions(authStore.permissions, ['view:event'])}
		<EventList previousEvents={data.previousEvents} upcomingEvents={data.upcomingEvents} />
	{:else}
		<span>Du har ikke tilladelse til at se denne side...</span>
	{/if}
</DashboardLayout>
