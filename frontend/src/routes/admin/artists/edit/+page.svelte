<script lang="ts">
	import ArtistForm from './ArtistForm.svelte';

	let { data } = $props();
</script>

<main class="px-auto min-h-svh py-20">
	<ArtistForm bind:artist={data.artist} genres={data.genres} />
</main>
