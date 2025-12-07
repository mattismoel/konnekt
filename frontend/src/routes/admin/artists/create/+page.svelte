<script lang="ts">
	import ArtistForm from "$lib/features/artists/components/ArtistForm.svelte";
	import { createArtist } from "$lib/features/artists/artist.remote";
</script>

<main class="mx-responsive min-h-svh py-32">
	<ArtistForm variant="create" form={createArtist} />
</main>
