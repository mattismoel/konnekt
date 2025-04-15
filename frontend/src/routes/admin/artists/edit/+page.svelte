<script lang="ts">
	import type { z } from 'zod';

	import { createArtist, updateArtist, type artistFormSchema } from '$lib/features/artist/artist';
	import { page } from '$app/state';

	import ArtistForm from './ArtistForm.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { APIError } from '$lib/api';

	let { data } = $props();

	const submit = async (form: z.infer<typeof artistFormSchema>) => {
		const id = page.url.searchParams.get('id');
		const isEdit = id !== null;

		try {
			isEdit ? await updateArtist(parseInt(id), form) : createArtist(form);
			toaster.addToast(`Kunstner ${isEdit ? 'opdateret' : 'skabt'}.`);
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast(
					`Kunne ikke ${isEdit ? 'opdatere' : 'lave'} kunstner ${e.status}`,
					e.cause,
					'error'
				);
			}

			toaster.addToast(`Kunne ikke ${isEdit ? 'opdatere' : 'lave'} kunstner...`, '', 'error');
		}
	};
</script>

<main class="flex min-h-svh justify-center p-16">
	<ArtistForm artist={data.artist} genres={data.genres} onSubmit={submit} />
</main>
