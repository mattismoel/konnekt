<script lang="ts">
	import { ZodError, type z } from 'zod';
	import { goto } from '$app/navigation';
	import { toaster } from '$lib/toaster.svelte';

	import { updateEvent, type editEventForm } from '$lib/features/event/event';
	import EventForm from '../../EventForm.svelte';

	let { data } = $props();
	let loading = $state(false);

	let errors = $state<z.typeToFlattenedError<z.infer<typeof editEventForm>>>();

	const handleSubmit = async (form: z.infer<typeof editEventForm>) => {
		try {
			loading = true;
			await updateEvent(fetch, form, data.event.id);
			toaster.addToast('Event opdateret');
			goto('/admin/events');
			loading = false;
		} catch (e) {
			if (e instanceof ZodError) {
				errors = e.flatten();
				toaster.addToast('Kunne ikke redigere event', 'Noget gik galt...', 'error');
				loading = false;
				return;
			}
			toaster.addToast('Kunne ikke redigere event', 'Noget gik galt...', 'error');
			loading = false;
			throw e;
		}
	};
</script>

<main>
	<h1 class="mb-8 text-4xl font-bold">Redigér event</h1>
	<EventForm
		{loading}
		event={data.event}
		{errors}
		venues={data.venues}
		artists={data.artists}
		onSubmit={handleSubmit}
	/>
</main>
