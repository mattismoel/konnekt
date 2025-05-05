<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import { hasPermissions } from '$lib/features/auth/permission';
	import ArtistList from './ArtistList.svelte';
	import * as AdminHeader from '$lib/components/admin-header/index';
	import { authStore } from '$lib/auth.svelte';

	let { data } = $props();
</script>

<AdminHeader.Root>
	<AdminHeader.Title>Kunstnere</AdminHeader.Title>
	<AdminHeader.Description
		>Overblik over alle kunstnere, som er associerede med events.</AdminHeader.Description
	>
	<AdminHeader.Actions>
		<Button
			href="/admin/artists/create"
			disabled={!hasPermissions(authStore.permissions, ['edit:artist'])}
		>
			<PlusIcon />Tilføj
		</Button>
	</AdminHeader.Actions>
</AdminHeader.Root>

<main class="pt-16">
	{#if hasPermissions(authStore.permissions, ['view:artist'])}
		<section class="space-y-4">
			<ArtistList artists={data.artists} upcomingArtists={data.upcomingArtists} />
		</section>
	{:else}
		<span>Du har ikke tilladelse til at se kunstnere...</span>
	{/if}
</main>
