import { DATE_FORMAT } from '@/lib/time';

import { format } from 'date-fns';

import { earliestConcert } from '@/lib/features/event/concert';
import type { Event } from '@/lib/features/event/event';

import Fader from "@/lib/components/fader"

import { FaCalendarDay, FaMapPin, FaMusic, FaTicketAlt } from "react-icons/fa"
import LinkButton from '@/lib/components/ui/button/link-button';
import { eventGenres } from '../../artist/genre';

type Props = {
	event: Event;
	active?: boolean;
	prefix?: string;
};

const EventDetails = ({ event, active, prefix }: Props) => {
	let fromDate = earliestConcert(event.concerts)?.from;

	const genres = eventGenres(event)
	let genresString = genres.map((genre) => genre.name).join(', ')

	const locationUrl = new URL(
		`https://www.google.com/maps/search/?` +
		new URLSearchParams({
			api: '1',
			query: `${event.venue.name},${event.venue.city},${event.venue.countryCode}`
		})
	)

	return (
		<div className="relative isolate flex h-[calc((100svh/5)*4)] items-end overflow-hidden pb-8 sm:pb-16">
			<img
				src={event.imageUrl}
				alt="Event cover"
				className="absolute top-0 left-0 -z-10 h-full w-full object-cover brightness-50"
			/>

			<Fader className="absolute -z-10 h-1/2" direction="up" />

			<div className="px-auto flex w-full flex-col">
				{prefix && <span className="text-shadow-sm">{prefix}</span>}

				<h1 className="md:mb-8 mb-4 text-5xl md:text-7xl font-bold text-shadow-md">{event.title}</h1>

				<div className="text-text/75 flex flex-col gap-8 sm:flex-row text-shadow-sm">
					<div className="flex flex-1 flex-col gap-1">
						{fromDate && (
							<div className="flex items-center gap-4">
								<FaCalendarDay />
								<span className="line-clamp-1">{format(fromDate, DATE_FORMAT)}</span>
							</div>
						)}
						<div className="flex items-center gap-4">
							<FaMapPin />
							<a className="line-clamp-1" href={locationUrl.toString()}
							>{event.venue.name}, {event.venue.city} ({event.venue.countryCode})</a
							>
						</div>
						<div className="flex items-center gap-4">
							<FaMusic />
							<span className="line-clamp-1">{genresString}</span>
						</div>
					</div>

					<div className="flex flex-col justify-end gap-2">
						<LinkButton to={event.ticketUrl} className="w-full"><FaTicketAlt />Køb billet</LinkButton>
						{!active && (
							<LinkButton
								to="/events/$eventId"
								params={{ eventId: event.id.toString() }}
								variant="outline"
								className="w-full"
							>
								Læs mere
							</LinkButton>
						)}
					</div>
				</div>
			</div>
		</div>
	)
}

export default EventDetails
