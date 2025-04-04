<script lang="ts">
	import { DATE_FORMAT } from '$lib/time';

	import type { Event } from '$lib/event';

	import { earliestConcert, latestConcert } from '$lib/concert';
	import { format } from 'date-fns';

	import TrashIcon from '~icons/mdi/trash';
	import Button from '$lib/components/ui/Button.svelte';

	type Props = {
		event: Event;
	};

	let { event }: Props = $props();

	const fromDate = $derived(earliestConcert(event.concerts)?.from || new Date());
	const toDate = $derived(latestConcert(event.concerts)?.to || new Date());
</script>

<li>
	<a
		class="flex items-center gap-4 rounded-md border border-transparent px-4 py-2 hover:border-zinc-800 hover:bg-zinc-900"
		href="/admin/events/edit?id={event.id}"
	>
		<span class="flex-1 font-medium">{event.title}</span>
		<span class="text-text/50 line-clamp-1 flex-1">{format(fromDate, DATE_FORMAT)}</span>
		<span class="text-text/50 line-clamp-1 flex-1"
			>{format(fromDate, 'HH:mm')} - {format(toDate, 'HH:mm')}</span
		>
		<Button variant="dangerous">
			<TrashIcon />
		</Button>
	</a>
</li>
