<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/ui/Button.svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import SearchBar from '$lib/components/ui/SearchBar.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import ArtistEntry from './ArtistEntry.svelte';

	let { data } = $props();

	let search = $state('');

	let artists = $derived(
		data.artists.filter((a) => a.name.toLowerCase().includes(search.toLowerCase()))
	);

	const deleteArtist = (id: number) => {};
</script>

<main class="space-y-8 px-16 py-16">
	<div>
		<div class="flex justify-between">
			<h1 class="font-heading mb-4 text-4xl font-bold">Kunstnere</h1>
			<Button onclick={() => goto('/admin/artists/edit')}><PlusIcon />Tilføj</Button>
		</div>
		<p class="text-text/50">Overblik over alle kunstnere, som er associerede med events.</p>
	</div>

	<section class="space-y-8">
		<SearchBar bind:value={search} />
		<ul>
			{#each artists as artist (artist.id)}
				<ArtistEntry {artist} onDelete={() => deleteArtist(artist.id)} />
			{/each}
		</ul>
	</section>
</main>
