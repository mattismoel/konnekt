<script lang="ts">
	import { APIError } from '$lib/api';
	import { deleteArtist, type Artist } from '$lib/features/artist/artist';
	import type { Permission } from '$lib/features/auth/permission';
	import { toaster } from '$lib/toaster.svelte';
	import ArtistEntry from './ArtistEntry.svelte';

	type Props = {
		artists: Artist[];
		memberPermissions: Permission[];
	};

	let { artists, memberPermissions }: Props = $props();

	const handleDeleteArtist = async (id: number) => {
		const artist = artists.find((a) => a.id === id);
		if (!artist) return;

		if (!confirm(`Vil du slette kunstner "${artist.name}"?`)) return;

		try {
			await deleteArtist(fetch, artist.id);
			toaster.addToast('Kunstner slettet');
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke slette kunstner', e.cause, 'error');
				return;
			}
			toaster.addToast('Kunne ikke slette kunstner', 'Noget gik galt', 'error');
			return;
		}
	};
</script>

<ul>
	{#each artists as artist (artist.id)}
		<ArtistEntry {artist} {memberPermissions} onDelete={() => handleDeleteArtist(artist.id)} />
	{/each}
</ul>
