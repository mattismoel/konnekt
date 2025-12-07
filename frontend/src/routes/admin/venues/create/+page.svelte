<script lang="ts">
	import VenueForm from "$lib/features/venues/components/VenueForm.svelte";
	import { createVenue } from "$lib/features/venues/venue.remote";
</script>

<main class="mx-responsive flex min-h-svh items-center justify-center py-32">
	<VenueForm variant="create" form={createVenue} />
</main>
