<script>
	import EditEventForm from './EditEventForm.svelte';

	const { data } = $props();
	const { event } = $derived(data);
</script>

<main class="min-h-sub-nav px-auto">
	<EditEventForm {event} />
</main>
