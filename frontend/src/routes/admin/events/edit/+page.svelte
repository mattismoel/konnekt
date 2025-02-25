<script lang="ts">
	import { z } from 'zod';
	import EditEventForm from './EditEventForm.svelte';
	import { eventForm } from '$lib/event';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	const { data } = $props();
	const { event, artists, venues } = $derived(data);

	const submit = (e: z.infer<typeof eventForm>) => {
		fetch(`${PUBLIC_BACKEND_URL}/events`, {
			method: 'POST',
			credentials: 'include'
		});
	};
</script>

<main class="min-h-sub-nav px-auto py-20">
	<EditEventForm onSubmit={submit} {event} artists={artists || []} venues={venues || []} />
</main>
