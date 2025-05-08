<script lang="ts">
	import { ZodError, type z } from 'zod';
	import VenueForm from '../../VenueForm.svelte';
	import { editVenue, editVenueForm } from '$lib/features/venue/venue';
	import { toaster } from '$lib/toaster.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import { hasPermissions } from '$lib/features/auth/permission';
	import { authStore } from '$lib/auth.svelte';

	let { data } = $props();

	let errors = $state<z.inferFlattenedErrors<typeof editVenueForm>>();

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

<main class="flex flex-col items-center justify-center">
	<div>
		<h1 class="font-heading mb-8 text-4xl font-bold">Redigér venue</h1>
		<VenueForm
			disabled={!hasPermissions(authStore.permissions, ['edit:venue'])}
			venue={data.venue}
			{errors}
			onSubmit={handleSubmit}
		/>
	</div>
</main>
