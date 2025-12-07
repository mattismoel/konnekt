<script lang="ts">
	import type { Concert, Event } from "$lib/features/events/event";
	import { addHours, differenceInMinutes, format, startOfHour } from "date-fns";
	import Button from "../../../components/ui/Button.svelte";
	import { generateGoogleCalendarEventUrl } from "$lib/google-calendar";

	type Props = {
		event: Event;
	};

	let { event }: Props = $props();

	const calendarUrl = $derived(generateGoogleCalendarEventUrl(event));

	let startHourDate = $derived(
		event.concerts.length > 0 ? startOfHour(event.concerts[0].fromDate) : startOfHour(new Date())
	);

	let endHour = $derived(
		event.concerts.length > 0
			? addHours(startOfHour(event.concerts[event.concerts.length - 1].toDate), 1)
			: addHours(startOfHour(new Date()), 4)
	);

	let timeMarkers = $derived.by(() => {
		let markers: Date[] = [];
		let current = startHourDate;

		while (current <= endHour) {
			markers = [...markers, current];
			current = addHours(current, 1);
		}

		return markers;
	});

	let totalMinutes = $derived(Math.max(1, differenceInMinutes(endHour, startHourDate)));
</script>

{#snippet entry(concert: Concert)}
	{@const concertStartOffset = differenceInMinutes(concert.fromDate, startHourDate)}
	{@const concertDurationMinutes = differenceInMinutes(concert.toDate, concert.fromDate)}

	<a
		style:top="calc({concertStartOffset / totalMinutes} * 100%)"
		style:height="calc({concertDurationMinutes / totalMinutes} * 100% - 1px)"
		class="group absolute flex min-h-10 w-full justify-between overflow-hidden rounded-xl border border-t border-blue-800 bg-blue-950 p-4 text-sm transition-colors hover:border-blue-600 hover:bg-blue-900"
		href="/artists/{concert.artist.id}"
	>
		<p class="font-bold text-blue-200 group-hover:underline">
			{concert.artist.name}
		</p>
		<p class="text-blue-500">
			{format(concert.fromDate, "HH:mm")} - {format(concert.toDate, "HH:mm")}
		</p>
	</a>
{/snippet}

<div class="flex w-full flex-col gap-8">
	<div class="flex items-center-safe justify-between">
		{#if event.concerts.length > 0}
			<div>
				<h3 class="mb-2 text-xl font-bold text-text-light">
					Program for {event.title}
				</h3>
				<span>{format(event.concerts[0].fromDate, "EEEE, dd/MM/yyyy")}</span>
			</div>
		{/if}

		{#if calendarUrl}
			<Button
				class="hidden sm:flex"
				variant="secondary"
				href={calendarUrl.toString()}
				target="__blank"
			>
				<!-- <FaCalendarPlus /> -->
				Føj til kalender
			</Button>
		{/if}
	</div>

	<div class="overflow-y-scroll">
		<div class="grid h-full min-h-72 flex-1 gap-4 sm:grid-cols-[48px_1fr]">
			<!-- {/*  Timeline */} -->
			<div class="relative hidden sm:block">
				{#each timeMarkers as marker, i}
					{@const markerOffset = differenceInMinutes(marker, startHourDate)}

					<div
						style:top="calc({markerOffset / totalMinutes} * 100%)"
						class={[
							"absolute w-full flex-col gap-1 text-sm",
							i === timeMarkers.length - 1 ? "hidden" : "flex"
						]}
					>
						<div class="h-px w-full bg-zinc-800"></div>
						<span>{format(marker, "HH:mm")}</span>
					</div>
				{/each}
			</div>

			<!-- {/* Events */} -->
			<div class="relative">
				{#each event.concerts as concert}
					{@render entry(concert)}
				{/each}
			</div>
		</div>

		{#if calendarUrl}
			<Button
				class="flex w-full sm:hidden"
				variant="secondary"
				href={calendarUrl.toString()}
				target="__blank"
			>
				<!-- <FaCalendarPlus /> -->
				Føj til kalender
			</Button>
		{/if}
	</div>
</div>
