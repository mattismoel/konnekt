<script lang="ts">
	import PlusIcon from '~icons/mdi/plus';
	import Button from '$lib/components/ui/Button.svelte';

	import { goto } from '$app/navigation';
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
				disabled={!hasPermissions(data.member.permissions, ['edit:event'])}
				onclick={() => goto(`/admin/events/create`)}
			>
				<PlusIcon />Tilføj
			</Button>
		</HeaderActions>
	</DashboardHeader>

	{#if hasPermissions(data.member.permissions, ['view:event'])}
		<div class="space-y-8">
			<SearchBar bind:value={search} />
			{#if search.trim() !== ''}
				<section>
					<EventList events={filteredEvents} memberPermissions={data.member.permissions} />
				</section>
			{:else}
				<section>
					<EventList events={data.upcomingEvents} memberPermissions={data.member.permissions} />
				</section>
			{/if}

			{#if data.previousEvents.length > 0}
				<section>
					<details>
						<summary class="mb-4">Tidligere events ({data.previousEvents.length})</summary>
						<EventList events={data.previousEvents} memberPermissions={data.member.permissions} />
					</details>
				</section>
			{/if}
		</div>
	{:else}
		<span>Du har ikke tilladelse til at se denne side...</span>
	{/if}
</DashboardLayout>
