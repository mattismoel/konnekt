<script lang="ts">
	import { ZodError, type z } from 'zod';
	import VenueForm from '../../VenueForm.svelte';
	import { editVenue, editVenueForm } from '$lib/features/venue/venue';
	import { toaster } from '$lib/toaster.svelte';
	import { goto, invalidateAll } from '$app/navigation';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<typeof editVenueForm>>();

	$inspect(data.venue);

	const handleSubmit = async (form: z.infer<typeof editVenueForm>) => {
		try {
			await editVenue(fetch, data.venue.id, form);
			toaster.addToast('Venue redigeret.');
			await invalidateAll();
			await goto('/admin/venues');
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
	<VenueForm venue={data.venue} {errors} onSubmit={handleSubmit} />
</main>
