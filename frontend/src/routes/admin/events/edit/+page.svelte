<script lang="ts">
	import { z } from 'zod';
	import EventForm from './EventForm.svelte';
	import { createEvent, eventForm, updateEvent } from '$lib/event';
	import { page } from '$app/state';

	const { data } = $props();
	const { event, artists, venues } = $derived(data);

	const submit = async (form: z.infer<typeof eventForm>) => {
		const id = page.url.searchParams.get('id');

		id ? await updateEvent(form, parseInt(id)) : await createEvent(form);
	};
</script>

<main class="min-h-sub-nav px-auto py-20">
	<EventForm onSubmit={submit} {event} artists={artists || []} venues={venues || []} />
</main>
