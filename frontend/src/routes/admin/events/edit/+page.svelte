<script lang="ts">
	import { z, ZodError } from 'zod';
	import { error } from '@sveltejs/kit';

	import { createEvent, eventForm, updateEvent } from '$lib/event';
	import { page } from '$app/state';
	import { APIError } from '$lib/error';

	import EventForm from './EventForm.svelte';

	const { data } = $props();
	const { event, artists, venues } = $derived(data);

	const submit = async (form: z.infer<typeof eventForm>) => {
		try {
			const id = page.url.searchParams.get('id');

			id ? await updateEvent(form, parseInt(id)) : await createEvent(form);
		} catch (e) {
			if (e instanceof APIError) {
				return error(e.status, e.message);
			}

			if (e instanceof ZodError) {
				console.error(e.issues);
				throw e;
			}

			throw e;
		}
	};
</script>

<main class="min-h-sub-nav px-auto py-20">
	<EventForm onSubmit={submit} {event} artists={artists || []} venues={venues || []} />
</main>
