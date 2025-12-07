<script lang="ts">
	import ArtistForm from "$lib/features/artists/components/ArtistForm.svelte";
	import { editArtist, getArtist } from "$lib/features/artists/artist.remote";
	let { params } = $props();

	const artist = $derived(await getArtist(params.id));
</script>

<main class="mx-responsive py-32">
	<ArtistForm variant="edit" form={editArtist} {artist} />
</main>
