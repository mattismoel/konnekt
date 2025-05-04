<script lang="ts">
	import { createArtist, createArtistForm } from '$lib/features/artist/artist';
	import { ZodError, type z } from 'zod';
	import ArtistForm from '../ArtistForm.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { goto } from '$app/navigation';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof createArtistForm>>>();
	let loading = $state(false);

	const handleSubmit = async (form: z.infer<typeof createArtistForm>) => {
		try {
			loading = true;
			await createArtist(fetch, form);
			toaster.addToast('Kunstner skabt');
			await goto('/admin/artists');
			loading = false;
		} catch (e) {
			if (e instanceof ZodError) {
				toaster.addToast('Kunne ikke skabe kunstner', 'Ugyldig kunstnerdata', 'error');
				errors = e.flatten();
				loading = false;
				throw e;
			}

			toaster.addToast('Kunne ikke skabe kunstner', 'Noget gik galt...', 'error');
			loading = false;
			throw e;
		}
	};
</script>

<main class="px-auto py-16 pt-32">
	<ArtistForm {loading} genres={data.genres} onSubmit={handleSubmit} {errors} />
</main>
