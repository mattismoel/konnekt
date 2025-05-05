<script lang="ts">
	import PlusIcon from '~icons/mdi/plus';
	import Button from '$lib/components/ui/Button.svelte';

	import { hasPermissions } from '$lib/features/auth/permission';
	import EventList from './EventList.svelte';
	import { authStore } from '$lib/auth.svelte';
	import * as AdminHeader from '$lib/components/admin-header/index';

	let { data } = $props();
</script>

<AdminHeader.Root>
	<AdminHeader.Title>Events</AdminHeader.Title>
	<AdminHeader.Description>Overblik over alle events.</AdminHeader.Description>
	<AdminHeader.Actions>
		<Button
			href="/admin/events/create"
			disabled={!hasPermissions(authStore.permissions, ['edit:event'])}
		>
			<PlusIcon />Tilføj
		</Button>
	</AdminHeader.Actions>
</AdminHeader.Root>

<main class="pt-16">
	{#if hasPermissions(authStore.permissions, ['view:event'])}
		<EventList previousEvents={data.previousEvents} upcomingEvents={data.upcomingEvents} />
	{:else}
		<span>Du har ikke tilladelse til at se denne side...</span>
	{/if}
</main>
