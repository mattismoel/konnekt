<script lang="ts">
	import { invalidateAll } from '$app/navigation';

	import type { z } from 'zod';

	import { toaster } from '$lib/toaster.svelte.js';

	import {
		createVenue,
		editVenue,
		deleteVenue,
		venueForm,
		type Venue
	} from '$lib/features/venue/venue.js';

	import Button from '$lib/components/ui/Button.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import VenueEntry from './VenueEntry.svelte';
	import SearchBar from '$lib/components/SearchBar.svelte';
	import { hasPermissions } from '$lib/features/auth/permission';
	import { tryCatch } from '$lib/error';
	import { APIError } from '$lib/api';

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
</script>

<main class="space-y-8 px-8 py-16 md:px-16">
	<div class="flex flex-col justify-between gap-8 md:flex-row">
		<div>
			<h1 class="font-heading mb-4 text-4xl font-bold">Venues</h1>
			<span class="text-text/50">
				Overblik over alle venues, som er associerede med events for Konnekt.
			</span>
		</div>
		<Button disabled={!hasPermissions(data.permissions, ['edit:venue'])}>
			<PlusIcon />Tilføj
		</Button>
	</div>

	{#if hasPermissions(data.permissions, ['view:venue'])}
		<section class="space-y-4">
			<SearchBar bind:value={search} />
			<ul>
				{#each venues as venue (venue.id)}
					<VenueEntry
						userPermissions={data.permissions}
						initialValue={venue}
						onEdit={(form) => handleEditVenue(venue.id, form)}
						onDelete={() => handleDeleteVenue(venue.id)}
					/>
				{/each}
			</ul>
		</section>
	{:else}
		<span>Du har ikke tilladelse til at se venues...</span>
	{/if}
</main>
