<script>
	import EditEventForm from './EditEventForm.svelte';

	const { data } = $props();
	const { event, artists, venues } = $derived(data);
</script>

<main class="min-h-sub-nav px-auto py-20">
	<EditEventForm {event} artists={artists || []} venues={venues || []} />
</main>
