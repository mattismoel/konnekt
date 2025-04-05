<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/ui/Button.svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import SearchBar from '$lib/components/ui/SearchBar.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import ArtistEntry from './ArtistEntry.svelte';
	import { hasPermissions } from '$lib/auth';

	let { data } = $props();

	let search = $state('');

	let artists = $derived(
		data.artists.filter((a) => a.name.toLowerCase().includes(search.toLowerCase()))
	);

	const deleteArtist = (id: number) => {};
</script>

<main class="space-y-8 px-8 py-16 md:px-16">
	<div class="flex flex-col justify-between gap-8 md:flex-row">
		<div class="">
			<h1 class="font-heading mb-4 text-4xl font-bold">Kunstnere</h1>
			<p class="text-text/50">Overblik over alle kunstnere, som er associerede med events.</p>
		</div>
		<Button
			disabled={!hasPermissions(data.permissions, ['edit:artist'])}
			onclick={() => goto('/admin/artists/edit')}
		>
			<PlusIcon />Tilføj
		</Button>
	</div>

	{#if hasPermissions(data.permissions, ['view:artist'])}
		<section class="space-y-8">
			<SearchBar bind:value={search} />
			<ul>
				{#each artists as artist (artist.id)}
					<ArtistEntry
						userPermissions={data.permissions}
						{artist}
						onDelete={() => deleteArtist(artist.id)}
					/>
				{/each}
			</ul>
		</section>
	{:else}
		<span>Du har ikke tilladelse til at se kunstnere...</span>
	{/if}
</main>
