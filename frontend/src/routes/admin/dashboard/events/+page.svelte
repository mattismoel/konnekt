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
	import SearchBar from '$lib/components/ui/SearchBar.svelte';

	let { data } = $props();

	let search = $state('');

	let upcomingEvents = $derived(
		data.upcomingEvents.filter((v) => v.title.toLowerCase().includes(search.toLowerCase()))
	);
</script>

<Card class="space-y-8">
	<div>
		<div class="flex justify-between">
			<h1 class="font-heading mb-4 text-4xl font-bold">Kommende events</h1>
			<Button onclick={() => goto(`/admin/events/edit`)}><PlusIcon />Tilføj</Button>
		</div>
		<p class="text-text/50">Overblik over alle events.</p>
	</div>

	<section class="space-y-4">
		<SearchBar bind:value={search} />
		<ul>
			{#each upcomingEvents as event (event.id)}
				<EventEntry {event} />
			{/each}
		</ul>
	</section>
</Card>
