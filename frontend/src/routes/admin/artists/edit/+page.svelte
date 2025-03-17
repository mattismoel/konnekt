<script lang="ts">
	import type { z } from 'zod';

	import { createArtist, updateArtist, type artistFormSchema } from '$lib/artist';
	import { page } from '$app/state';

	import ArtistForm from './ArtistForm.svelte';

	let { data } = $props();

	const submit = async (form: z.infer<typeof artistFormSchema>) => {
		const id = page.url.searchParams.get('id');
		id ? await updateArtist(parseInt(id), form) : createArtist(form);
	};
</script>

<main class="flex min-h-svh justify-center py-20">
	<ArtistForm artist={data.artist} genres={data.genres} onSubmit={submit} />
</main>
