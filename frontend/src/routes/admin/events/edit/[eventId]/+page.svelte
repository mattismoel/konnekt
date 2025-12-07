<script lang="ts">
	import { page } from "$app/state";
	import EventForm from "$lib/features/events/components/EventForm.svelte";
	import { editEvent, getEvent } from "$lib/features/events/event.remote";
</script>

<main class="mx-responsive py-32">
	<EventForm form={editEvent} variant="edit" event={await getEvent(page.params.eventId ?? "")} />
</main>
