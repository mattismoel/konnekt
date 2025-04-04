<script lang="ts">
	import { invalidateAll } from '$app/navigation';

	import type { z } from 'zod';

	import { toaster } from '$lib/toaster.svelte.js';
	import { APIError, tryCatch } from '$lib/error.js';

	import { COUNTRIES_MAP } from '$lib/location.js';
	import { createVenue, editVenue, deleteVenue, venueForm, type Venue } from '$lib/venue.js';

	import Button from '$lib/components/ui/Button.svelte';
	import Selector from '$lib/components/ui/Selector.svelte';

	import Table from '$lib/components/ui/table/Table.svelte';
	import TableBody from '$lib/components/ui/table/TableBody.svelte';
	import TableCell from '$lib/components/ui/table/TableCell.svelte';
	import TableHead from '$lib/components/ui/table/TableHead.svelte';
	import TableHeader from '$lib/components/ui/table/TableHeader.svelte';
	import TableRow from '$lib/components/ui/table/TableRow.svelte';
	import InputCell from '$lib/components/ui/table/InputCell.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import TrashIcon from '~icons/mdi/trash';
	import EditIcon from '~icons/mdi/edit';
	import Card from '$lib/components/ui/Card.svelte';
	import VenueEntry from './VenueEntry.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import SearchBar from '$lib/components/ui/SearchBar.svelte';

	let { data } = $props();

	let search = $state('');

	let venues = $derived(
		data.venues.filter((v) => v.name.toLowerCase().includes(search.toLowerCase()))
	);

	let form = $state<z.infer<typeof venueForm>>({
		name: '',
		city: '',
		countryCode: 'DK'
	});

	const handleEditVenue = async (id: number, form: z.infer<typeof venueForm>) => {
		const { error } = await tryCatch(editVenue(id, form));
		if (error) {
			if (error instanceof APIError) {
				toaster.addToast('Kunne ikke regigere venue', error.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke redigere venue', 'Noget gik galt...', 'error');
			return;
		}

		toaster.addToast('Venue redigeret');
	};

	const handleAddVenue = async (form: z.infer<typeof venueForm>) => {
		const { error } = await tryCatch(createVenue(form));
		if (error) {
			if (error instanceof APIError) {
				toaster.addToast('Kunne ikke tilføje venue', error.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke tilføje venue', 'Noget gik galt...', 'error');
			return;
		}

		toaster.addToast('Venue tilføjet');
		await invalidateAll();
	};

	const handleDeleteVenue = async (id: number) => {
		const venue = venues.find((v) => v.id === id);
		if (!venue) return;

		if (!confirm(`Er sikker på, at du vil slette venue "${venue.name}"?`)) return;

		const { error } = await tryCatch(deleteVenue(id));
		if (error) {
			if (error instanceof APIError) {
				toaster.addToast('Kunne ikke slette venue', error.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke slette venue', 'Noget gik galt...', 'error');
			return;
		}

		toaster.addToast('Venue slettet');
		await invalidateAll();
	};

	let selectedVenue = $state<Venue | undefined>(undefined);
	let showVenueModal = $state(false);

	const toggleVenueModal = (venueId: number) => {
		const venue = venues.find((v) => v.id === venueId);

		selectedVenue = venue;

		showVenueModal = true;
	};

	$inspect(selectedVenue);
</script>

<main class="space-y-8 px-16 py-16">
	<div>
		<div class="flex justify-between">
			<h1 class="font-heading mb-4 text-4xl font-bold">Venues</h1>
			<Button><PlusIcon />Tilføj</Button>
		</div>
		<span class="text-text/50">
			Overblik over alle venues, som er associerede med events for Konnekt.
		</span>
	</div>

	<section class="space-y-4">
		<SearchBar bind:value={search} />
		<ul>
			{#each venues as venue (venue.id)}
				<VenueEntry
					initialValue={venue}
					onEdit={(form) => handleEditVenue(venue.id, form)}
					onDelete={() => handleDeleteVenue(venue.id)}
				/>
			{/each}
		</ul>
	</section>
</main>
