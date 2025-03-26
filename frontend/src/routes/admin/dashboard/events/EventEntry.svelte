<script lang="ts">
	import { DATETIME_FORMAT } from '$lib/time';

	import type { Event } from '$lib/event';

	import { earliestConcert, latestConcert } from '$lib/concert';
	import { format } from 'date-fns';

	import TrashIcon from '~icons/mdi/trash';

	type Props = {
		event: Event;
	};

	let { event }: Props = $props();

	const fromDate = $derived(earliestConcert(event.concerts)?.from || new Date());
	const toDate = $derived(latestConcert(event.concerts)?.to || new Date());
</script>

<a
	class="flex items-center gap-4 rounded-md border border-transparent p-2 hover:border-zinc-800 hover:bg-zinc-900"
	href="/admin/events/edit?id={event.id}"
>
	<span class="flex-1">{event.title}</span>
	<span class="line-clamp-1 flex-1 text-zinc-500">{format(fromDate, DATETIME_FORMAT)}</span>
	<span class="line-clamp-1 flex-1 text-zinc-500">{format(toDate, DATETIME_FORMAT)}</span>
	<button class="text-zinc-400 hover:text-red-500">
		<TrashIcon />
	</button>
</a>
