<script lang="ts">
	import { z } from 'zod';
	import { error } from '@sveltejs/kit';

	import { createEvent, eventForm, updateEvent } from '$lib/features/event/event';
	import { page } from '$app/state';
	import { APIError } from '$lib/api';

	import EventForm from './EventForm.svelte';
	import { toaster } from '$lib/toaster.svelte';

	const { data } = $props();
	const { event, artists, venues } = $derived(data);

	const submit = async (form: z.infer<typeof eventForm>) => {
		const id = page.url.searchParams.get('id');
		const isEdit = id !== null;

		try {
			isEdit ? await updateEvent(form, parseInt(id)) : await createEvent(form);
			toaster.addToast(`Event ${isEdit ? 'opdateret' : 'skabt'}.`);
		} catch (e) {
			if (e instanceof APIError) {
				const msg = `Kunne ikke ${isEdit ? 'opdatere' : 'skabe'} event (${e.status})`;
				toaster.addToast(msg, e.cause, 'error');
				return error(e.status, e.message);
			}

			toaster.addToast(`Kunne ikke ${isEdit ? 'opdatere' : 'skabe'} event`);

			throw e;
		}
	};
</script>

<main class="min-h-sub-nav px-16 py-20">
	<EventForm onSubmit={submit} {event} artists={artists || []} venues={venues || []} />
</main>
