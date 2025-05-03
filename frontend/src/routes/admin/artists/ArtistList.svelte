<script lang="ts">
	import List from '$lib/components/ui/List.svelte';
	import { type Artist } from '$lib/features/artist/artist';
	import type { Permission } from '$lib/features/auth/permission';
	import ArtistEntry from './ArtistEntry.svelte';

	type Props = {
		artists: Artist[];
		memberPermissions: Permission[];
	};

	let { artists, memberPermissions }: Props = $props();
</script>

<List>
	{#each artists as artist (artist.id)}
		<ArtistEntry {artist} {memberPermissions} />
	{/each}
</List>
