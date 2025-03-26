<script lang="ts">
	import { invalidateAll } from '$app/navigation';

	import type { z } from 'zod';

	import { toaster } from '$lib/toaster.svelte.js';
	import { APIError } from '$lib/error.js';

	import { COUNTRIES_MAP } from '$lib/location.js';
	import { createVenue, deleteVenue, venueForm } from '$lib/venue.js';

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

	let { data } = $props();
	let { venues } = $derived(data);

	const form = $state<z.infer<typeof venueForm>>({
		city: '',
		countryCode: 'DK',
		name: ''
	});

	const handleAddVenue = async () => {
		try {
			await createVenue(form);
			toaster.addToast('Venue tilføjet');

			await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke tilføje venue', e.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke tilføje venue', 'Noget gik galt...', 'error');
		}
	};

	const handleDeleteVenue = async (id: number) => {
		if (!venues.some((v) => v.id === id)) return;

		const venue = venues.reduce((prev, curr) => (curr.id === id ? curr : prev));

		if (!confirm(`Er sikker på, at du vil slette venue "${venue.name}"?`)) return;

		try {
			await deleteVenue(id);
			toaster.addToast('Venue slettet');
			await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke slette venue', e.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke slette venue', 'Noget gik galt...', 'error');
		}
	};
</script>

<main class="px-auto min-h-svh py-20">
	<h1 class="mb-4 text-4xl font-bold">Venues</h1>

	<Table class="mb-8">
		<TableHead>
			<TableRow>
				<TableHeader>Venuenavn</TableHeader>
				<TableHeader>By</TableHeader>
				<TableHeader>Land</TableHeader>
				<TableHeader alignment="right">Handling</TableHeader>
			</TableRow>
		</TableHead>
		<TableBody>
			<TableRow class="hover:bg-zinc-950">
				<InputCell bind:value={form.name} placeholder="Venuenavn..." />
				<InputCell bind:value={form.city} placeholder="By" />
				<TableCell>
					<Selector
						class="h-min w-full"
						entries={Array.from(COUNTRIES_MAP).map(([key, value]) => ({ name: value, value: key }))}
						bind:value={form.countryCode}
					/>
				</TableCell>
				<TableCell alignment="right">
					<div class="flex w-full justify-end">
						<Button variant="ghost" onclick={handleAddVenue}><PlusIcon /> Tilføj</Button>
					</div>
				</TableCell>
			</TableRow>
			{#each venues as venue (venue.id)}
				<TableRow>
					<TableCell>{venue.name}</TableCell>
					<TableCell class="text-text/50">{venue.city}</TableCell>
					<TableCell class="text-text/50">
						{COUNTRIES_MAP.get(venue.countryCode)}, {venue.countryCode}
					</TableCell>
					<TableCell alignment="right">
						<div class="flex justify-end gap-2">
							<Button variant="dangerous" class="w-full" onclick={() => handleDeleteVenue(venue.id)}
								><TrashIcon /></Button
							>
						</div>
					</TableCell>
				</TableRow>
			{/each}
		</TableBody>
	</Table>
</main>
