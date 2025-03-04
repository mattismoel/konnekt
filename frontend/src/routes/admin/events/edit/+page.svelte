<script lang="ts">
	import { z } from 'zod';
	import EventForm from './EventForm.svelte';
	import { createEvent, eventForm } from '$lib/event';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import { page } from '$app/state';

	const { data } = $props();
	const { event, artists, venues } = $derived(data);

	const submit = async (form: z.infer<typeof eventForm>) => {
		const id = page.url.searchParams.get('id');

		if (!id) {
			const event = await createEvent(form);
		} else {
			// TODO: Implement updateEvent() function.
			const event = await updateEvent(form);
		}

		fetch(`${PUBLIC_BACKEND_URL}/events`, {
			method: 'POST',
			credentials: 'include'
		});
	};
</script>

<main class="min-h-sub-nav px-auto py-20">
	<EventForm onSubmit={submit} {event} artists={artists || []} venues={venues || []} />
</main>
