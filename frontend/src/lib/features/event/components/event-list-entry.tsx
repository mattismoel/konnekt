import { useState } from "react";
import { format, isBefore, startOfToday } from "date-fns";

import { APIError } from "@/lib/api";
import { DATETIME_FORMAT } from "@/lib/time";
import { useToast } from "@/lib/context/toast";
import { useAuth } from "@/lib/context/auth";
import { cn } from "@/lib/clsx";

import { earliestConcert } from "../concert";
import { deleteEvent, type Event } from "../event";
import type { Artist } from "../../artist/artist";

import List from "@/lib/components/list/list";
import ContextMenu from "@/lib/components/context-menu";

import { FaMapMarker } from "react-icons/fa";

type Props = {
	event: Event;
};

const formatArtists = (artists: Artist[]): string => {
	if (artists.length > 2) {
		return artists.slice(0, 2).map(a => a.name).join(",") + " " + `(+${artists.length - 2} mere)`
	}

	return artists.map(a => a.name).join(", ")
}

const EventEntry = ({ event }: Props) => {
	const { addToast } = useToast()
	const { hasPermissions } = useAuth()
	let [showContextMenu, setShowContextMenu] = useState(false);

	let artists = event.concerts.map((concert) => concert.artist);

	const fromDate = earliestConcert(event.concerts)?.from || new Date()
	let expired = isBefore(fromDate, startOfToday());

	const handleDeleteEvent = async () => {
		if (!confirm(`Vil du slette ${event.title}?`)) return;

		try {
			await deleteEvent(event.id);
			addToast('Event slettet');
			// await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) {
				addToast('Kunne ikke slette event', e.cause, 'error');
				return;
			}

			addToast('Kunne ikke slette event', 'Noget gik galt...', 'error');
			return;
		}
	};

	return (
		<List.Entry title="Redigér event" className={`group ${expired ? 'expired' : ''}`}>
			<List.Entry.LinkSection to="/admin/events/$eventId/edit" params={{ eventId: event.id.toString() }}>
				<span className="line-clamp-1 group-[.expired]:line-through">{event.title}</span>
				<span className="text-text/50 line-clamp-1">{format(fromDate, DATETIME_FORMAT)}</span>
				<span className="text-text/50 line-clamp-1 md:hidden">{event.venue.name}</span>
				<span className="text-text/50 line-clamp-1 hidden md:block">
					{formatArtists(artists)}
				</span>
			</List.Entry.LinkSection>
			<List.Entry.LinkSection to="/admin/venues/$venueId/edit" params={{ venueId: event.venue.id.toString() }} className="group/venue hidden md:block">
				<span
					// className:disabled={!hasPermissions(authStore.permissions, ['edit:venue'])}
					className={cn("text-text/50 group-hover/venue:text-text group-[.disabled]/venue:text-text/50 hidden w-full items-center gap-2 group-hover/venue:underline group-[.disabled]/venue:no-underline md:flex", {
						"disabled": hasPermissions(["edit:venue"])
					})}
				>
					<FaMapMarker />
					<span className="whitespace-nowrap">{event.venue.name}</span>
				</span>
			</List.Entry.LinkSection>

			<List.Entry.Section className='w-min'>
				<ContextMenu.Button onClick={() => setShowContextMenu(true)} />
			</List.Entry.Section>

			<ContextMenu show={showContextMenu} onClose={() => setShowContextMenu(false)} className="absolute top-1/2 right-4">
				<ContextMenu.LinkEntry
					disabled={!hasPermissions(['edit:event'])}
					to="/admin/events/$eventId/edit"
					params={{ eventId: event.id.toString() }}
				>
					Redigér
				</ContextMenu.LinkEntry>
				<ContextMenu.Entry
					disabled={!hasPermissions(['delete:event'])}
					onClick={handleDeleteEvent}
				>
					Slet
				</ContextMenu.Entry>
			</ContextMenu>
		</List.Entry>
	)
}

export default EventEntry
