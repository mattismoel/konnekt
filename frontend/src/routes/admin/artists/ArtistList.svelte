<script lang="ts">
	import SearchList from '$lib/components/SearchList.svelte';
	import { type Artist } from '$lib/features/artist/artist';
	import type { Permission } from '$lib/features/auth/permission';
	import ArtistEntry from './ArtistEntry.svelte';

	type Props = {
		artists: Artist[];
		upcomingArtists: Artist[];
	};

	let { artists, upcomingArtists }: Props = $props();

	let search = $state('');

	const filteredArtists = $derived(
		artists.filter((a) => a.name.toLowerCase().includes(search.toLowerCase()))
	);
</script>

<SearchList bind:search>
	{#if search}
		{#each filteredArtists as artist (artist.id)}
			<ArtistEntry {artist} />
		{/each}
	{:else}
		{#each upcomingArtists as artist (artist.id)}
			<ArtistEntry {artist} />
		{/each}

		<details>
			<summary class="mb-4">Alle kunstnere</summary>
			{#each artists as artist (artist.id)}
				<ArtistEntry {artist} />
			{/each}
		</details>
	{/if}
</SearchList>
