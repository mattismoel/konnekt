<script lang="ts">
	import Card from '$lib/components/ui/Card.svelte';
	import Table from '$lib/components/ui/table/Table.svelte';
	import TableBody from '$lib/components/ui/table/TableBody.svelte';
	import TableCell from '$lib/components/ui/table/TableCell.svelte';
	import TableHead from '$lib/components/ui/table/TableHead.svelte';
	import TableHeader from '$lib/components/ui/table/TableHeader.svelte';
	import TableRow from '$lib/components/ui/table/TableRow.svelte';
	import { earliestConcert, latestConcert } from '$lib/concert';
	import { format } from 'date-fns';
	import EventEntry from './EventEntry.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import { DATE_FORMAT, DATETIME_FORMAT } from '$lib/time';
	import Button from '$lib/components/ui/Button.svelte';

	import TrashIcon from '~icons/mdi/trash';
	import CleanIcon from '~icons/mdi/broom';

	import EditIcon from '~icons/mdi/edit';
	import { goto, invalidateAll } from '$app/navigation';
	import SearchBar from '$lib/components/SearchBar.svelte';
	import { hasPermissions } from '$lib/features/auth/permission';
	import EventList from './EventList.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { deleteEvent } from '$lib/features/event/event';
	import { tryCatch } from '$lib/error';
	import { APIError } from '$lib/api';

	let { data } = $props();

	let search = $state('');

	let filteredEvents = $derived(
		[...data.previousEvents, ...data.upcomingEvents].filter((event) =>
			event.title.toLowerCase().includes(search.toLowerCase())
		)
	);

	let listPreviousEvents = $state(false);

	const handleCleanPreviousEvents = async () => {
		let message = 'Er du sikker på, at du vil rydde alle tidligere events?\n\n';

		data.previousEvents.map((e) => {
			message += `- ${e.title}\n`;
		});

		message += '\n';
		message += 'Handlingen kan ikke fortrydes.';

		if (!confirm(message)) return;

		data.previousEvents.forEach(async ({ id }) => {
			const { error } = await tryCatch(deleteEvent(fetch, id));
			if (error) {
				if (error instanceof APIError) {
					toaster.addToast('Kunne ikke rydde op', error.cause, 'error');
					return;
				}

				toaster.addToast('Kunne ikke rydde op', 'Noget gik galt...', 'error');
				return;
			}
		});

		toaster.addToast('Events blev ryddet op');
		await invalidateAll();
	};

	const handleDeleteEvent = async (id: number) => {
		const event = [...data.upcomingEvents, ...data.previousEvents].find((event) => event.id === id);

		if (!event) return;

		if (!confirm(`Vil du slette ${event.title}?`)) return;

		const { error } = await tryCatch(deleteEvent(fetch, id));
		if (error) {
			if (error instanceof APIError) {
				toaster.addToast('Kunne ikke slette event', error.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke slette event', 'Noget gik galt...', 'error');
			return;
		}

		toaster.addToast('Event slettet');
		await invalidateAll();
	};
</script>

<main class="space-y-8 px-8 py-16 md:px-16">
	<div class="flex flex-col justify-between gap-8 md:flex-row">
		<div>
			<h1 class="font-heading mb-4 text-4xl font-bold md:line-clamp-1">Events</h1>
			<p class="text-text/50">Overblik over alle events.</p>
		</div>
		<div class="flex gap-2">
			<Button
				disabled={!hasPermissions(data.permissions, ['edit:event'])}
				onclick={() => goto(`/admin/events/edit`)}
			>
				<PlusIcon />Tilføj
			</Button>
			<Button onclick={handleCleanPreviousEvents} variant="ghost"><CleanIcon />Ryd</Button>
		</div>
	</div>

	{#if hasPermissions(data.permissions, ['view:event'])}
		<div class="space-y-8">
			<SearchBar bind:value={search} />
			{#if search.trim() !== ''}
				<section>
					<EventList onDelete={handleDeleteEvent} events={filteredEvents} />
				</section>
			{:else}
				<section>
					<EventList onDelete={handleDeleteEvent} events={data.upcomingEvents} />
				</section>
			{/if}

			{#if data.previousEvents.length > 0}
				<section>
					<details>
						<summary class="mb-4">Tidligere events ({data.previousEvents.length})</summary>
						<EventList events={data.previousEvents} onDelete={handleDeleteEvent} />
					</details>
				</section>
			{/if}
		</div>
	{:else}
		<span>Du har ikke tilladelse til at se denne side...</span>
	{/if}
</main>
