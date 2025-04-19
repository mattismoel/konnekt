<script lang="ts">
	import { ZodError, type z } from 'zod';
	import EventForm from '../EventForm.svelte';
	import { goto } from '$app/navigation';
	import { toaster } from '$lib/toaster.svelte';

	import { updateEvent, type editEventForm } from '$lib/features/event/event';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof editEventForm>>>();

	const handleSubmit = async (form: z.infer<typeof editEventForm>) => {
		try {
			await updateEvent(fetch, form, data.eventId);
			toaster.addToast('Event opdateret');
			goto('/admin/dashboard/events');
		} catch (e) {
			if (e instanceof ZodError) {
				errors = e.flatten();
				toaster.addToast('Kunne ikke redigere event', 'Noget gik galt...', 'error');
				return;
			}
			toaster.addToast('Kunne ikke redigere event', 'Noget gik galt...', 'error');
			throw e;
		}
	};
</script>

<main class="px-8 py-16 md:px-16">
	<h1 class="mb-8 text-4xl font-bold">Redigér event</h1>
	<EventForm
		event={data.event}
		{errors}
		venues={data.venues}
		artists={data.artists}
		onSubmit={handleSubmit}
	/>
</main>
