<script lang="ts">
	import PlusIcon from '~icons/mdi/plus';
	import Button from '$lib/components/ui/Button.svelte';

	import SearchBar from '$lib/components/SearchBar.svelte';
	import { hasPermissions } from '$lib/features/auth/permission';
	import EventList from './EventList.svelte';
	import DashboardLayout from '../DashboardLayout.svelte';
	import DashboardHeader from '../DashboardHeader.svelte';
	import HeaderActions from '../HeaderActions.svelte';

	let { data } = $props();

	let search = $state('');

	let filteredEvents = $derived(
		[...data.previousEvents, ...data.upcomingEvents].filter((event) =>
			event.title.toLowerCase().includes(search.toLowerCase())
		)
	);
</script>

<DashboardLayout>
	<DashboardHeader title="Events" description="Overblik over alle events.">
		<HeaderActions>
			<Button
				href="/admin/events/create"
				disabled={!hasPermissions(data.member.permissions, ['edit:event'])}
			>
				<PlusIcon />Tilføj
			</Button>
		</HeaderActions>
	</DashboardHeader>

	{#if hasPermissions(data.member.permissions, ['view:event'])}
		<EventList
			previousEvents={data.previousEvents}
			upcomingEvents={data.upcomingEvents}
			memberPermissions={data.member.permissions}
		/>
	{:else}
		<span>Du har ikke tilladelse til at se denne side...</span>
	{/if}
</DashboardLayout>
