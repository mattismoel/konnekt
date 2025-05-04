<script lang="ts">
	import { createEvent, createEventForm } from '$lib/features/event/event';
	import { z, ZodError } from 'zod';
	import EventForm from '../EventForm.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { goto } from '$app/navigation';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof createEventForm>> | undefined>();
	let loading = $state(false);

	const handleSubmit = async (form: z.infer<typeof createEventForm>) => {
		try {
			loading = true;
			await createEvent(fetch, form);
			toaster.addToast('Event skabt', 'Event blev skabt fejlfrit');
			goto('/admin/events');
			loading = false;
		} catch (e) {
			if (e instanceof ZodError) {
				errors = e.flatten();
				loading = false;
				return;
			}
			loading = false;
			toaster.addToast('Event kunne ikke skabes', 'Der skete en fejl', 'error');
			throw e;
		}
	};
</script>

<main class="px-auto py-16 pt-32">
	<EventForm
		{loading}
		venues={data.venues}
		artists={data.artists}
		{errors}
		onSubmit={handleSubmit}
	/>
</main>
