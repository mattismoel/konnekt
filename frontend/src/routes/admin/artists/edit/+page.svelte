<script lang="ts">
	import { toaster } from '$lib/toaster.svelte.js';
	import { ZodError, type z } from 'zod';
	import ArtistForm from '../ArtistForm.svelte';
	import { updateArtist, type editArtistForm } from '$lib/features/artist/artist';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof editArtistForm>>>();

	const handleSubmit = async (form: z.infer<typeof editArtistForm>) => {
		try {
			await updateArtist(fetch, data.artist.id, form);
			toaster.addToast('Kunstner opdateret');
		} catch (e) {
			if (e instanceof ZodError) {
				toaster.addToast('Kunne ikke opdatere kunstner', 'Ugyldig kunstnerdata', 'error');
				errors = e.flatten();
				return;
			}

			toaster.addToast('Kunne ikke opdatere kunstner', 'Noget gik galt...', 'error');
			throw e;
		}
	};
</script>

<main class="flex min-h-svh justify-center p-16">
	<ArtistForm {errors} artist={data.artist} genres={data.genres} onSubmit={handleSubmit} />
</main>
