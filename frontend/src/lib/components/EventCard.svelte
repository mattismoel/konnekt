<script lang="ts">
	import type { Event } from '$lib/event';
	import { formatDateStr } from '$lib/time';
	import CalendarIcon from '~icons/mdi/calendar';
	import VenueIcon from '~icons/mdi/map-marker';
	import QRCode from 'qrcode';

	type Props = {
		event: Event;
	};

	const { event }: Props = $props();
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
	href={`/events/${event.id}`}
	class="group relative overflow-hidden rounded-md bg-red-500 bg-gradient-to-tr from-zinc-950 to-zinc-900 p-[1px] transition-colors duration-700 hover:to-zinc-700"
>
	<div
		class="absolute top-0 right-32 h-12 w-12 -translate-y-1/2 translate-x-1/2 rounded-full border border-zinc-900 bg-black"
	></div>
	<div
		class="absolute right-32 bottom-0 h-12 w-12 translate-x-1/2 translate-y-1/2 rounded-full border border-zinc-900 bg-black"
	></div>
	<div
		class="group flex h-40 overflow-hidden rounded-md bg-gradient-to-t from-zinc-950 via-zinc-900 via-80% to-zinc-950 transition-colors group-hover:via-zinc-800"
	>
		<!-- Image -->
		<img src={event.coverImageUrl} alt={`Cover for ${event.title}`} class="w-40 object-cover" />
		<!-- Information -->
		<div
			class="absolute z-50 h-full w-full bg-black opacity-30 transition-opacity duration-500 group-hover:opacity-0"
		></div>
		<div class="flex w-56 flex-col justify-between border-r-2 border-dashed border-zinc-700 p-4">
			<span class="text-xl font-bold text-zinc-200">{event.title}</span>
			<div class="space-y-1 text-sm text-zinc-300">
				<div class="flex gap-2">
					<span><CalendarIcon /></span>
					<span>{formatDateStr(earliestConcert.from)}</span>
				</div>
				<div class="flex gap-2">
					<span><VenueIcon /></span>
					<span>{event.venue.name}, {event.venue.city}</span>
				</div>
			</div>
			<div class="flex flex-col">
				<span class="text-xs text-zinc-500"><b>Billetnr.:</b> {ticketCode}</span>
				<span class="font-black text-zinc-400">KONNEKT</span>
			</div>
		</div>
		<!-- QR -->
		<div class="*: flex w-32 flex-col items-center justify-center gap-2 px-8 py-2">
			<span class="text-sm font-bold text-zinc-300">SCAN</span>
			<canvas bind:this={qrCodeCanvas}></canvas>
			<div class="flex flex-col items-center text-xs text-zinc-600">
				<span class="font-bold">Billetnr.:</span>
				<span>{ticketCode}</span>
			</div>
		</div>
	</div>
</a>
