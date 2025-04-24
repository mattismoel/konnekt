<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/ui/Button.svelte';
	import SearchBar from '$lib/components/SearchBar.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import { hasPermissions } from '$lib/features/auth/permission';
	import ArtistList from './ArtistList.svelte';
	import DashboardLayout from '../DashboardLayout.svelte';
	import DashboardHeader from '../DashboardHeader.svelte';
	import HeaderActions from '../HeaderActions.svelte';

	let { data } = $props();

	let search = $state('');

	let artists = $derived(
		data.artists.filter((a) => a.name.toLowerCase().includes(search.toLowerCase()))
	);
</script>

<DashboardLayout>
	<DashboardHeader
		title="Kunstnere"
		description="Overblik over alle kunstnere, som er associerede med events."
	>
		<HeaderActions>
			<Button
				disabled={!hasPermissions(data.member.permissions, ['edit:artist'])}
				onclick={() => goto('/admin/artists/create')}
			>
				<PlusIcon />Tilføj
			</Button>
		</HeaderActions>
	</DashboardHeader>

	{#if hasPermissions(data.member.permissions, ['view:artist'])}
		<section class="space-y-4">
			<SearchBar bind:value={search} />
			<ArtistList {artists} memberPermissions={data.member.permissions} />
		</section>
	{:else}
		<span>Du har ikke tilladelse til at se kunstnere...</span>
	{/if}
</DashboardLayout>

<!-- <main class="space-y-8 px-8 py-16 md:px-16"> -->
<!-- 	<div class="flex flex-col justify-between gap-8 md:flex-row"> -->
<!-- 		<div class=""> -->
<!-- 			<h1 class="font-heading mb-4 text-4xl font-bold">Kunstnere</h1> -->
<!-- 			<p class="text-text/50">Overblik over alle kunstnere, som er associerede med events.</p> -->
<!-- 		</div> -->
<!-- 		<Button -->
<!-- 			disabled={!hasPermissions(data.member.permissions, ['edit:artist'])} -->
<!-- 			onclick={() => goto('/admin/artists/create')} -->
<!-- 		> -->
<!-- 			<PlusIcon />Tilføj -->
<!-- 		</Button> -->
<!-- 	</div> -->
<!---->
<!-- 	{#if hasPermissions(data.member.permissions, ['view:artist'])} -->
<!-- 		<section class="space-y-8"> -->
<!-- 			<SearchBar bind:value={search} /> -->
<!-- 		</section> -->
<!-- 		<ArtistList {artists} memberPermissions={data.member.permissions} /> -->
<!-- 	{:else} -->
<!-- 		<span>Du har ikke tilladelse til at se kunstnere...</span> -->
<!-- 	{/if} -->
<!-- </main> -->
