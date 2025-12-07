<script lang="ts">
	import VenueForm from "$lib/features/venues/components/VenueForm.svelte";
	import { editVenue, getVenue } from "$lib/features/venues/venue.remote";

	let { params } = $props();
	const venue = $derived(await getVenue(params.id));
</script>

<main class="mx-responsive flex min-h-svh items-center justify-center py-32">
	<VenueForm variant="edit" form={editVenue} {venue} />
</main>
