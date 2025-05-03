<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import { hasPermissions } from '$lib/features/auth/permission';
	import ArtistList from './ArtistList.svelte';
	import DashboardLayout from '../DashboardLayout.svelte';
	import DashboardHeader from '../DashboardHeader.svelte';
	import HeaderActions from '../HeaderActions.svelte';

	let { data } = $props();
</script>

<DashboardLayout>
	<DashboardHeader
		title="Kunstnere"
		description="Overblik over alle kunstnere, som er associerede med events."
	>
		<HeaderActions>
			<Button
				href="/admin/artists/create"
				disabled={!hasPermissions(data.member.permissions, ['edit:artist'])}
			>
				<PlusIcon />Tilføj
			</Button>
		</HeaderActions>
	</DashboardHeader>

	{#if hasPermissions(data.member.permissions, ['view:artist'])}
		<section class="space-y-4">
			<ArtistList
				artists={data.artists}
				upcomingArtists={data.upcomingArtists}
				memberPermissions={data.member.permissions}
			/>
		</section>
	{:else}
		<span>Du har ikke tilladelse til at se kunstnere...</span>
	{/if}
</DashboardLayout>
