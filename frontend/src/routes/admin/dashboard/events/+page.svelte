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
	import EditIcon from '~icons/mdi/edit';
	import { goto } from '$app/navigation';

	let { data } = $props();

	let { upcomingEvents } = $derived(data);
</script>

<Card>
	<div class="flex justify-between">
		<h1 class="font-heading mb-4 text-4xl font-bold">Kommende events</h1>
		<a href="/admin/events/edit" class="flex gap-2">
			<PlusIcon class="text-xl" />
			Tilføj event
		</a>
	</div>
	<section>
		<Table>
			<TableHead>
				<TableRow>
					<TableHeader>Eventtitel</TableHeader>
					<TableHeader>Dato</TableHeader>
					<TableHeader>Tidspunkt</TableHeader>
					<TableHeader></TableHeader>
				</TableRow>
			</TableHead>
			<TableBody class="divide-zinc-800">
				{#each upcomingEvents as event (event.id)}
					{@const fromDate = earliestConcert(event.concerts)?.from || new Date()}
					{@const toDate = latestConcert(event.concerts)?.to || new Date()}

					<TableRow class="hover:bg-zinc-800">
						<TableCell class="font-medium">{event.title}</TableCell>
						<TableCell class="text-text/50">{format(fromDate, DATE_FORMAT)}</TableCell>
						<TableCell class="text-text/50"
							>{format(fromDate, 'HH:mm')}-{format(toDate, 'HH:mm')}</TableCell
						>
						<TableCell>
							<div class="flex justify-end gap-6">
								<button
									title="Redigér event"
									onclick={() => goto(`/admin/events/edit?id=${event.id}`)}
									class="text-text/50 hover:text-text"
								>
									<EditIcon />
								</button>
								<Button variant="dangerous">
									<TrashIcon />
								</Button>
							</div>
						</TableCell>
					</TableRow>
				{/each}
			</TableBody>
		</Table>
	</section></Card
>
