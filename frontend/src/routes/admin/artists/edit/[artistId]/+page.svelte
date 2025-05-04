<script lang="ts">
	import { toaster } from '$lib/toaster.svelte.js';
	import { ZodError, type z } from 'zod';
	import { updateArtist, type editArtistForm } from '$lib/features/artist/artist';
	import { goto } from '$app/navigation';
	import ArtistForm from '../../ArtistForm.svelte';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof editArtistForm>>>();

	let loading = $state(false);

	const handleSubmit = async (form: z.infer<typeof editArtistForm>) => {
		try {
			loading = true;
			await updateArtist(fetch, data.artist.id, form);
			toaster.addToast('Kunstner opdateret');
			loading = false;
			goto('/admin/artists');
		} catch (e) {
			if (e instanceof ZodError) {
				toaster.addToast('Kunne ikke opdatere kunstner', 'Ugyldig kunstnerdata', 'error');
				errors = e.flatten();
				loading = false;
				return;
			}

			loading = false;
			toaster.addToast('Kunne ikke opdatere kunstner', 'Noget gik galt...', 'error');
			throw e;
		}
	};
</script>

<main class="px-auto py-16 pt-32">
	<ArtistForm
		{loading}
		{errors}
		artist={data.artist}
		genres={data.genres}
		onSubmit={handleSubmit}
	/>
</main>
