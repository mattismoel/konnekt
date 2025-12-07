<script lang="ts">
	import EventForm from "$lib/features/events/components/EventForm.svelte";
	import { createEvent } from "$lib/features/events/event.remote";
</script>

<main class="mx-responsive py-32">
	<EventForm variant="create" form={createEvent} />
</main>
