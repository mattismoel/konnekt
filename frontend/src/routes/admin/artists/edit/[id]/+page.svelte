<script lang="ts">
	import { page } from "$app/state";
	import ArtistForm from "$lib/features/artists/components/ArtistForm.svelte";
	import { editArtist, getArtist } from "$lib/features/artists/artist.remote";

	const artist = await getArtist(page.params.id!);
</script>

<main class="mx-responsive py-32">
	<ArtistForm variant="edit" form={editArtist} {artist} />
</main>
