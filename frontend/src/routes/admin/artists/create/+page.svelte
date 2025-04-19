<script lang="ts">
	import { createArtist, createArtistForm } from '$lib/features/artist/artist';
	import { ZodError, type z } from 'zod';
	import ArtistForm from '../ArtistForm.svelte';
	import { toaster } from '$lib/toaster.svelte';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof createArtistForm>>>();

	const handleSubmit = async (form: z.infer<typeof createArtistForm>) => {
		try {
			await createArtist(fetch, form);
			toaster.addToast("Kunstner skabt")
		} catch (e) {
			if (e instanceof ZodError) {
				toaster.addToast('Kunne ikke skabe kunstner', 'Ugyldig kunstnerdata', 'error');
				errors = e.flatten();
				return;
			}
		}
	};
</script>

<main class="px-8 py-16 md:px-16">
	<ArtistForm genres={data.genres} onSubmit={handleSubmit} {errors} />
</main>
