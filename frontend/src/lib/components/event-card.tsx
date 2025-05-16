import type { Event } from '@/lib/features/event';
// import CalendarIcon from '~icons/mdi/calendar';
// import VenueIcon from '~icons/mdi/map-marker';
import QRCode from 'qrcode';
// import type { HTMLAttributes } from 'svelte/elements';
import { cn } from '@/lib/clsx';
import Logo from '@/lib/assets/logo';
import { format } from 'date-fns';
import { DATETIME_FORMAT } from '@/lib/time';

import { MdCalendarToday, MdMap } from 'react-icons/md';
import { useEffect, useRef, type HTMLAttributes } from 'react';
import { Link } from '@tanstack/react-router';

type Props = HTMLAttributes<HTMLDivElement> & {
	event: Event;
};


const EventCard = ({ event, className }: Props) => {
	let ticketCode = Math.floor(10e9 * Math.random());
	let canvas = useRef<HTMLCanvasElement>(null);

	const earliestConcert = event.concerts[0];

	useEffect(() => {
		if (!canvas.current) return

		QRCode.toCanvas(canvas.current, event.ticketUrl, {
			margin: 1,
			width: 64
		}, (e) => {
			if (e) console.error(e)
		});
	}, [canvas])

	return (
		<Link
			to="/events/$eventId"
			params={{ eventId: event.id.toString() }}
			aria-labelledby="title"
			className="group relative isolate shrink-0 overflow-hidden"
		>
			<Holes />
			{/* {@render holes()} */}
			<div
				className="absolute z-50 h-full w-full bg-zinc-950 opacity-0 transition-opacity duration-200 ease-out group-hover:opacity-0 md:opacity-30"
			></div>
			<div
				className={cn(
					'h-40 rounded-md bg-gradient-to-tr from-zinc-900 to-zinc-800 p-[1px] transition-colors duration-500 group-hover:to-zinc-700',
					className
				)}
			>
				<div
					className="zinc-900 flex h-full shrink-0 overflow-hidden rounded-md bg-gradient-to-t from-zinc-950 via-zinc-900 via-80% to-zinc-950"
				>
					<img
						src={event.imageUrl}
						alt="Cover for {event.title}"
						className="aspect-square h-full object-cover"
					/>

					{/* <!-- INFORMATION --> */}
					<div
						className="flex h-full w-full min-w-fit flex-col border-r-[2px] border-dashed border-zinc-700 p-4"
					>
						<h3 className="text-xl font-bold text-zinc-300">{event.title}</h3>
						<div className="flex flex-1 flex-col justify-center text-zinc-400">
							<div className="flex gap-2">
								<MdCalendarToday />
								{/* <CalendarIcon /> */}
								<time>{format(earliestConcert.from, DATETIME_FORMAT)}</time>
							</div>
							<div className="flex gap-2">
								<MdMap />
								<address className="not-italic">{event.venue.name}, {event.venue.city}</address>
							</div>
						</div>
						<div className="flex flex-col gap-1 text-xs text-zinc-500">
							<span><b>Billetnr:</b> {ticketCode}</span>
							<Logo />
						</div>
					</div>
					<div className="flex h-full w-28 shrink-0 flex-col items-center justify-center gap-1 p-3 text-xs">
						<span><b>SCAN</b></span>
						<canvas ref={canvas}></canvas>
						<span className="text-center"><b>Billetnr:</b><br />{ticketCode}</span>
					</div>
				</div>
			</div>
		</Link>

	)
}

const Holes = () => {
	return (
		<>
			<div
				className="absolute top-0 right-28 z-10 h-12 w-12 translate-x-1/2 -translate-y-1/2 rounded-full bg-zinc-800 bg-gradient-to-r from-zinc-700 to-zinc-900 p-[1px]"
			>
				<div className="h-full w-full rounded-full bg-zinc-950"></div>
			</div>
			<div
				className="absolute right-28 bottom-0 z-10 h-12 w-12 translate-x-1/2 translate-y-1/2 rounded-full bg-gradient-to-r from-zinc-800 to-zinc-900 p-[1px]"
			>
				<div className="h-full w-full rounded-full bg-zinc-950"></div>
			</div>
		</>
	)
}

export default EventCard

