import { useState } from "react";
import type { Event } from "../event";
import { format } from "date-fns";
import { earliestConcert } from "../concert";
import { DATETIME_FORMAT } from "@/lib/time";
import { Link } from "@tanstack/react-router";
import Fader from "@/lib/components/fader";

type Props = {
	event: Event;
};

const EventCard = ({ event }: Props) => {
	const fromDate = earliestConcert(event.concerts)?.from

	type Pos = {
		x: number;
		y: number;
	};

	let [mousePos, setMousePos] = useState<Pos>({ x: 0, y: 0 });

	return (

		<Link
			className="group w-full"
			to="/events/$eventId"
			params={{ eventId: event.id.toString() }}
		>
			<div
				role="none"
				className="relative h-64 w-full overflow-hidden"
				onMouseMove={(e) => {
					const rect = e.currentTarget.getBoundingClientRect();
					setMousePos(() => ({
						x: e.clientX - rect.left,
						y: e.clientY - rect.top
					}))
				}}
			>
				<img
					src={event.imageUrl}
					alt={event.title}
					className="h-full w-full scale-110 object-cover transition-all duration-200 group-hover:scale-100 group-hover:brightness-100 md:brightness-90"
				/>
				<Fader direction="up" className="absolute h-48 from-black/80" />
				<div className="absolute bottom-0 left-0 h-full w-full border 
					border-white/0 mix-blend-overlay transition-all 
					group-hover:border-white/50"
				/>
				<div className="absolute bottom-0 left-0 flex flex-col px-5 pb-5 
					text-text transition-all duration-100 md:translate-y-full 
					md:group-hover:translate-y-0"
				>
					<h3 className="mb-2 text-3xl font-bold">{event.title}</h3>
					<div className="text-text/75 flex flex-col">
						{fromDate && (
							<span>{format(fromDate, DATETIME_FORMAT)}</span>
						)}
						<span>{event.venue.name}, {event.venue.city}</span>
					</div>
				</div>
				<div
					className="pointer-events-none absolute h-72 w-72 -translate-x-1/2 
					-translate-y-1/2 scale-0 bg-white/50 mix-blend-overlay  blur-3xl 
					transition-transform duration-400 group-hover:scale-100"
					style={{
						left: `${mousePos.x}px`,
						top: `${mousePos.y}px`
					}}
				></div>
			</div>
		</Link >
	)
}

export default EventCard
