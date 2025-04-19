<script lang="ts">
	import { createEvent, createEventForm } from '$lib/features/event/event';
	import { z, ZodError } from 'zod';
	import EventForm from '../EventForm.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { goto } from '$app/navigation';

	let { data } = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof createEventForm>> | undefined>();

	const handleSubmit = async (form: z.infer<typeof createEventForm>) => {
		try {
			await createEvent(fetch, form);
			toaster.addToast('Event skabt', 'Event blev skabt fejlfrit');
		} catch (e) {
			if (e instanceof ZodError) {
				errors = e.flatten();
			}
			toaster.addToast('Event kunne ikke skabes', 'Der skete en fejl', 'error');
			goto('/admin/dashboard/events');
			throw e;
		}
	};
</script>

<main class="px-8 py-16 md:px-16">
	<EventForm venues={data.venues} artists={data.artists} {errors} onSubmit={handleSubmit} />
</main>
