<script lang="ts">
	import { earliestConcert, type Event } from "$lib/features/events/event";
	import { DATETIME_FORMAT } from "$lib/time";
	import { format } from "date-fns";

	type Props = { event: Event };
	let { event }: Props = $props();
	const fromDate = $derived(earliestConcert(event.concerts)?.fromDate);
	let mousePos = $state<{ x: number; y: number }>({ x: 0, y: 0 });
</script>

<a class="group w-full" href="/events/{event.id}">
	<div
		role="none"
		class="relative h-64 w-full overflow-hidden rounded-xl"
		onmousemove={(e) => {
			const rect = e.currentTarget.getBoundingClientRect();
			mousePos = {
				x: e.clientX - rect.left,
				y: e.clientY - rect.top
			};
		}}
	>
		<img
			src={event.cover}
			alt={event.title}
			loading="lazy"
			class="h-full w-full scale-110 object-cover brightness-75 transition-all duration-200 group-hover:scale-100 group-hover:brightness-100 md:brightness-90"
		/>
		<!-- <Fader direction="up" class="absolute h-48 from-black/80" /> -->
		<div
			class="absolute bottom-0 left-0 h-full w-full border border-white/0 mix-blend-overlay transition-all group-hover:border-white/50"
		></div>
		<div
			class="absolute bottom-0 left-0 flex flex-col p-8 transition-all duration-100 md:translate-y-full md:group-hover:translate-y-0"
		>
			<span class="mb-2 text-3xl font-bold text-text-light text-shadow-md">
				{event.title}
			</span>
			<div class="flex flex-col text-base text-shadow-sm">
				{#if fromDate}
					<span>{format(fromDate, DATETIME_FORMAT)}</span>
				{/if}
				<span>
					{event.venue.name}, {event.venue.city}
				</span>
			</div>
		</div>
		<div
			class="pointer-events-none absolute h-72 w-72 -translate-x-1/2 -translate-y-1/2 scale-0 bg-white/50 mix-blend-overlay blur-3xl transition-transform duration-400 group-hover:scale-100"
			style:left="{mousePos.x}px"
			style:top="{mousePos.y}px"
		></div>
	</div>
</a>
