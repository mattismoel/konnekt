import { useRef } from "react";
import { format } from "date-fns"
import { Link } from "@tanstack/react-router";
import { useMousePos } from "@/lib/hooks/useMousePos";
import { DATE_FORMAT } from "@/lib/date";
import type { Event } from "../event";
import GlowingCursor from "@/lib/components/glowing-cursor";

type Props = {
	event: Event;
};

const EventCard = ({ event }: Props) => {
	const ref = useRef<HTMLDivElement>(null)

	const mousePos = useMousePos(ref)

	return (
		<Link to="/events/$eventId" params={{ eventId: event.id.toString() }}>
			<div ref={ref} className="group relative overflow-hidden h-64 rounded-sm">
				<img src={event.imageUrl} className="light-border-1 absolute top-0 left-0 w-full h-full object-cover brightness-80 scale-108 fade-b-16 to-black transition-[scale,filter] duration-500 group-hover:scale-100 group-hover:brightness-100" />

				{/* We need to explicitly add new elements for the image fade and border, as the ::after properties do not work on <img/> elements */}
				<div className="fade-to-b-56 fade-background opacity-75 transition-opacity group-hover:opacity-100 duration-500" />
				<div className="light-border-1 light-border-opacity-0 after:duration-1000 group-hover:light-border-opacity-75" />

				<GlowingCursor mousePos={mousePos} className="transition-[scale,opacity] not-group-hover:scale-50 not-group-hover:opacity-0" />

				<div className="absolute left-8 bottom-8 md:-bottom-6 transition-[bottom] duration-200 md:group-hover:bottom-8" >
					<h3 className="text-4xl text-heading font-semibold font-heading mb-2 transition-colors md:text-text md:group-hover:text-heading">{event.title}</h3>
					<div className="flex flex-col transition-opacity md:opacity-0 md:group-hover:opacity-100">
						<span>{format(event.concerts[0].from, DATE_FORMAT)}</span>
						<span>{event.venue.name}, {event.venue.city} ({event.venue.country})</span>
					</div>
				</div>
			</div>
		</Link>
	)
}

export default EventCard
