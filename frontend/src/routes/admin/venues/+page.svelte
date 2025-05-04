<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import { hasPermissions } from '$lib/features/auth/permission';
	import VenueList from './VenueList.svelte';
	import { authStore } from '$lib/auth.svelte';
	import * as AdminHeader from '$lib/components/admin-header/index';

	let { data } = $props();
</script>

<AdminHeader.Root>
	<AdminHeader.Title>Venues</AdminHeader.Title>
	<AdminHeader.Description
		>Overblik over alle venues, som er associerede med events for Konnekt.</AdminHeader.Description
	>
	<AdminHeader.Actions>
		<Button
			href="/admin/venues/create"
			disabled={!hasPermissions(authStore.permissions, ['edit:venue'])}
		>
			<PlusIcon />Tilføj
		</Button>
	</AdminHeader.Actions>
</AdminHeader.Root>

<main class="pt-16">
	{#if hasPermissions(authStore.permissions, ['view:venue'])}
		<VenueList venues={data.venues} />
	{:else}
		<span>Du har ikke tilladelse til at se venues...</span>
	{/if}
</main>
