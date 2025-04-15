<script lang="ts">
	import type { Artist } from '$lib/features/artist/artist';

	import Modal from '$lib/components/ui/modal/Modal.svelte';
	import ArtistSelectEntry from './ArtistSelectEntry.svelte';

	type Props = {
		artists: Artist[];
		show: boolean;
		onClose: () => void;
		onSelect: (artistId: number) => void;
	};

	let { artists, show, onSelect, onClose }: Props = $props();

	let selectedArtist: Artist | null = $state(null);

	const selectArtist = (id: number) => {
		selectedArtist = artists.reduce((curr, prev) => (curr.id === id ? curr : prev));
	};
</script>

<Modal title="Vælg kunstner..." {show} {onSelect} {onClose}>
	{#each artists as artist (artist.id)}
		<ArtistSelectEntry
			{artist}
			onSelect={() => selectArtist(artist.id)}
			onDeselect={() => (selectedArtist = null)}
			selected={selectedArtist?.id === artist.id}
		/>
	{/each}
</Modal>
