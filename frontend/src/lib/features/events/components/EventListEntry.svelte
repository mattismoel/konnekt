<script lang="ts">
	import type { Artist } from "$lib/features/artist";
	import type { Event } from "$lib/features/event";
	import { DATETIME_FORMAT } from "$lib/time";
	import { format } from "date-fns";

	type Props = {
		event: Event;
	};

	let { event }: Props = $props();

	const artistNames = $derived(event.concerts.map((concert) => concert.artist.name));
</script>

<li
	class="group flex items-center rounded-xl border border-zinc-800 bg-zinc-900 transition-colors hover:border-zinc-700 hover:bg-zinc-800"
>
	<a href="/admin/events/edit/{event.id}" class="w-full py-4 pl-8">
		<div>
			<p class="font-medium transition-colors group-hover:text-text-light group-hover:underline">
				{event.title}
			</p>
			<div class="text-base">
				<p>{format(event.concerts[0].fromDate, DATETIME_FORMAT)}</p>
				<p>
					{#if artistNames.length > 2}
						{artistNames.slice(0, 2).join(", ")} (+{artistNames.length - 2})
					{:else}
						{artistNames.join(", ")}
					{/if}
				</p>
			</div>
		</div>
	</a>
</li>
