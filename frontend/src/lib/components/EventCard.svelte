<script lang="ts">
	import type { Event } from '$lib/event';
	import { formatDateStr } from '$lib/time';
	import CalendarIcon from '~icons/mdi/calendar';
	import VenueIcon from '~icons/mdi/map-marker';
	import QRCode from 'qrcode';
	import type { HTMLAttributes } from 'svelte/elements';
	import { cn } from '$lib/clsx';
	import Logo from '$lib/assets/Logo.svelte';

	type Props = HTMLAttributes<HTMLDivElement> & {
		event: Event;
	};

	const { event, ...rest }: Props = $props();
	const earliestConcert = $derived(event.concerts[0]);

	let ticketCode = Math.floor(10e9 * Math.random());

	let qrCodeCanvas: HTMLCanvasElement;

	$effect(() => {
		QRCode.toCanvas(qrCodeCanvas, `https://konnekt.dk/events/${event.id}`, {
			margin: 1,
			width: 64
		});
	});
</script>

<a
	href="/events/{event.id}"
	aria-labelledby="title"
	class="group relative isolate shrink-0 overflow-hidden"
>
	{@render holes()}
	<div
		class="absolute z-50 h-full w-full bg-zinc-950 opacity-0 transition-opacity duration-300 ease-out group-hover:opacity-0 md:opacity-30"
	></div>
	<div
		class={cn('h-40 rounded-md bg-gradient-to-tr from-zinc-900 to-zinc-700 p-[1px]', rest.class)}
	>
		<div
			class="zinc-900 flex h-full shrink-0 overflow-hidden rounded-md bg-gradient-to-t from-zinc-950 via-zinc-900 via-80% to-zinc-950"
		>
			<img
				src={event.imageUrl}
				alt="Cover for {event.title}"
				class="aspect-square h-full object-cover"
			/>

			<!-- INFORMATION -->
			<div
				class="flex h-full w-full min-w-fit flex-col border-r-[2px] border-dashed border-zinc-700 p-4"
			>
				<h3 class="text-xl font-bold text-zinc-300">{event.title}</h3>
				<div class="flex flex-1 flex-col justify-center text-zinc-400">
					<div class="flex gap-2">
						<CalendarIcon />
						<time>{formatDateStr(earliestConcert.from)}</time>
					</div>
					<div class="flex gap-2">
						<VenueIcon />
						<address class="not-italic">{event.venue.name}, {event.venue.city}</address>
					</div>
				</div>
				<div class="flex flex-col gap-1 text-xs text-zinc-500">
					<span><b>Billetnr:</b> {ticketCode}</span>
					<Logo />
				</div>
			</div>
			<div class="flex h-full w-28 shrink-0 flex-col items-center justify-center gap-1 p-3 text-xs">
				<!-- QR-CODE -->
				<span><b>SCAN</b></span>
				<canvas bind:this={qrCodeCanvas}></canvas>
				<span class="text-center"><b>Billetnr:</b><br />{ticketCode}</span>
			</div>
		</div>
	</div>
</a>

{#snippet holes()}
	<div
		class="absolute top-0 right-28 z-10 h-12 w-12 -translate-y-1/2 translate-x-1/2 rounded-full bg-zinc-800 bg-gradient-to-r from-zinc-700 to-zinc-900 p-[1px]"
	>
		<div class="h-full w-full rounded-full bg-zinc-950"></div>
	</div>
	<div
		class="absolute right-28 bottom-0 z-10 h-12 w-12 translate-x-1/2 translate-y-1/2 rounded-full bg-gradient-to-r from-zinc-800 to-zinc-900 p-[1px]"
	>
		<div class="h-full w-full rounded-full bg-zinc-950"></div>
	</div>
{/snippet}
