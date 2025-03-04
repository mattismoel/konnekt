<script lang="ts">
	import type { z } from 'zod';
	import ArtistForm from './ArtistForm.svelte';
	import { createArtist, updateArtist, type artistFormSchema } from '$lib/artist';
	import { page } from '$app/state';
	import CreateConcertCard from '../../events/edit/CreateConcertCard.svelte';

	let { data } = $props();

	const submit = async (form: z.infer<typeof artistFormSchema>) => {
		const id = page.url.searchParams.get('id');
		if (id) {
			const artist = await updateArtist(parseInt(id), form);
			console.log('updated', artist);
			return;
		}

		const artist = await createArtist(form);
		console.log('created', artist);
	};
</script>

<main class="flex min-h-svh justify-center py-20">
	<ArtistForm artist={data.artist} genres={data.genres} onSubmit={submit} />
</main>
