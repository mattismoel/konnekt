<script lang="ts">
	import { earliestConcert, eventGenres, type Event } from "$lib/features/events/event";
	import { format } from "date-fns";
	import Button from "../../../components/ui/Button.svelte";
	import { DATE_FORMAT } from "$lib/time";

	type Props = {
		event: Event;
		active?: boolean;
		prefix?: string;
	};

	let { event, prefix, active }: Props = $props();
	let fromDate = $derived(earliestConcert(event.concerts)?.fromDate);

	const genres = $derived(
		eventGenres(event)
			.slice(0, 5)
			.map((genre) => genre.name)
			.join(", ")
	);

	const locationUrl = $derived(
		new URL(
			`https://www.google.com/maps/search/?` +
				new URLSearchParams({
					api: "1",
					query: `${event.venue.name},${event.venue.city},${event.venue.countryCode}`
				})
		)
	);
</script>

<div class="relative isolate flex h-[calc((100svh/5)*4)] items-end overflow-hidden pb-8 sm:pb-16">
	<img
		src={event.cover}
		alt="Event cover"
		class="absolute top-0 left-0 -z-10 h-full w-full object-cover brightness-50"
	/>

	<!-- <Fader class="absolute -z-10 h-1/2" direction="up" /> -->

	<div class="mx-responsive flex w-full flex-col">
		{#if prefix}
			<span class="text-shadow-sm">{prefix}</span>
		{/if}

		<h1
			class="font-heading mb-6 text-5xl font-bold text-text-light text-shadow-md md:mb-8 md:text-7xl"
		>
			{event.title}
		</h1>

		<div class="flex flex-col gap-8 sm:flex-row">
			<div class="flex flex-1 flex-col justify-end gap-2 text-base">
				{#if fromDate}
					<div class="flex items-center gap-4">
						<!-- <FaCalendarDay /> -->
						<span class="line-clamp-1">
							{format(fromDate, DATE_FORMAT)}
						</span>
					</div>
				{/if}
				<div class="flex items-center gap-4">
					<!-- <FaMusic /> -->
					<span class="line-clamp-1">{genres}</span>
				</div>
				<div class="flex items-center gap-4">
					<!-- <FaMapPin /> -->
					<a class="line-clamp-1 hover:underline" href={locationUrl.toString()}>
						{event.venue.name} / {event.venue.city} (
						{event.venue.countryCode})
					</a>
				</div>
			</div>

			<div class="flex flex-col justify-end gap-2">
				<Button href={event.ticketUrl} class="w-full">
					<!-- <FaTicketAlt /> -->
					Find billeter
				</Button>
				{#if !active}
					<Button href="/events/{event.id}" variant="secondary" class="w-full">Læs mere</Button>
				{/if}
			</div>
		</div>
	</div>
</div>
