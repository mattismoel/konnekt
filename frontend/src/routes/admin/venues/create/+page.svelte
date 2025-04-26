<script lang="ts">
	import { ZodError, type z } from 'zod';
	import VenueForm from '../VenueForm.svelte';
	import { createVenue, createVenueForm } from '$lib/features/venue/venue';
	import { toaster } from '$lib/toaster.svelte';
	import { goto, invalidateAll } from '$app/navigation';

	let errors = $state<z.typeToFlattenedError<typeof createVenueForm>>();

	const handleSubmit = async (form: z.infer<typeof createVenueForm>) => {
		try {
			await createVenue(fetch, form);
			toaster.addToast('Venue skabt.');
			await invalidateAll();
			await goto('/admin/dashboard/venues');
		} catch (e) {
			if (e instanceof ZodError) {
				errors = e.flatten();
				return;
			}
			toaster.addToast('Kunne ikke lave venue', 'Noget gik galt...', 'error');
		}
	};
</script>

<main class="flex items-center justify-center px-8 py-16 md:px-16">
	<VenueForm {errors} onSubmit={handleSubmit} />
</main>
