<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/ui/Button.svelte';
	import SearchBar from '$lib/components/SearchBar.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import ArtistEntry from './ArtistEntry.svelte';
	import { hasPermissions } from '$lib/features/auth/permission';
	import { toaster } from '$lib/toaster.svelte';
	import { deleteArtist } from '$lib/features/artist/artist';
	import { APIError } from '$lib/api';
	import { tryCatch } from '$lib/error';

	let { data } = $props();

	let search = $state('');

	let artists = $derived(
		data.artists.filter((a) => a.name.toLowerCase().includes(search.toLowerCase()))
	);

	const handleDeleteArtist = async (id: number) => {
		const artist = artists.find((a) => a.id === id);
		if (!artist) return;

		if (!confirm(`Vil du slette kunstner "${artist.name}"?`)) return;

		const { error } = await tryCatch(deleteArtist(fetch, artist.id));

		if (error) {
			if (error instanceof APIError) {
				toaster.addToast('Kunne ikke slette kunstner', error.cause, 'error');
				return;
			}
			toaster.addToast('Kunne ikke slette kunstner', 'Noget gik galt', 'error');
			return;
		}

		toaster.addToast('Kunstner slettet');
	};
</script>

<main class="space-y-8 px-8 py-16 md:px-16">
	<div class="flex flex-col justify-between gap-8 md:flex-row">
		<div class="">
			<h1 class="font-heading mb-4 text-4xl font-bold">Kunstnere</h1>
			<p class="text-text/50">Overblik over alle kunstnere, som er associerede med events.</p>
		</div>
		<Button
			disabled={!hasPermissions(data.member.permissions, ['edit:artist'])}
			onclick={() => goto('/admin/artists/create')}
		>
			<PlusIcon />Tilføj
		</Button>
	</div>

	{#if hasPermissions(data.member.permissions, ['view:artist'])}
		<section class="space-y-8">
			<SearchBar bind:value={search} />
			<ul>
				{#each artists as artist (artist.id)}
					<ArtistEntry
						{artist}
						memberPermissions={data.member.permissions}
						onDelete={() => handleDeleteArtist(artist.id)}
					/>
				{/each}
			</ul>
		</section>
	{:else}
		<span>Du har ikke tilladelse til at se kunstnere...</span>
	{/if}
</main>
