<script lang="ts">
	import SearchBar from '$lib/components/SearchBar.svelte';
	import List from '$lib/components/ui/List.svelte';
	import { type Artist } from '$lib/features/artist/artist';
	import type { Permission } from '$lib/features/auth/permission';
	import ArtistEntry from './ArtistEntry.svelte';

	type Props = {
		artists: Artist[];
		upcomingArtists: Artist[];
		memberPermissions: Permission[];
	};

	let { artists, upcomingArtists, memberPermissions }: Props = $props();

	let search = $state('');

	const filteredArtists = $derived(
		artists.filter((a) => a.name.toLowerCase().includes(search.toLowerCase()))
	);
</script>

<div class="flex flex-col gap-8">
	<SearchBar bind:value={search} />
	{#if search}
		<List>
			{#each filteredArtists as artist (artist.id)}
				<ArtistEntry {artist} {memberPermissions} />
			{/each}
		</List>
	{:else}
		<List>
			{#each upcomingArtists as artist (artist.id)}
				<ArtistEntry {artist} {memberPermissions} />
			{/each}
		</List>

		<details>
			<summary class="mb-4">Alle kunstnere</summary>
			<List>
				{#each artists as artist (artist.id)}
					<ArtistEntry {artist} {memberPermissions} />
				{/each}
			</List>
		</details>
	{/if}
</div>
