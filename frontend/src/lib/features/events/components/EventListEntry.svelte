<script lang="ts">
	import type { Event } from "$lib/features/events/event";
	import { DATETIME_FORMAT } from "$lib/time";
	import { format, isBefore, startOfToday } from "date-fns";

	type Props = {
		event: Event;
	};

	let { event }: Props = $props();
</script>

<li class="flex flex-col divide-y divide-zinc-800 rounded-2xl border border-zinc-800 bg-zinc-900">
	<div class="px-8 py-4">
		<a href="/admin/events/edit/{event.id}" class="hover:underline">
			<p
				class={[
					"mb-2 text-2xl font-bold text-text-light",

					isBefore(event.concerts[0].fromDate, startOfToday()) && "line-through"
				]}
			>
				{event.title}
			</p>
		</a>

		<a href="/admin/venues/edit/{event.venue.id}" class="hover:underline">
			{event.venue.name}
		</a>

		<p>{format(event.concerts[0].fromDate, DATETIME_FORMAT)}</p>
	</div>

	<div class="px-8 py-4">
		<ul class="flex gap-2">
			{#each event.concerts as concert}
				<li class="flex">
					<a
						href="/admin/artists/edit/{concert.artist.id}"
						class="group rounded-full border border-zinc-700 bg-zinc-800 px-6 py-2 text-sm transition-colors hover:border-zinc-600 hover:bg-zinc-700"
					>
						{concert.artist.name}
					</a>
				</li>
			{/each}
		</ul>
	</div>
</li>
