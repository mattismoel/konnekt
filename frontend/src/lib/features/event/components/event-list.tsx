import { useMemo, useState } from 'react';

import EventEntry from './event-list-entry';
import type { Event } from '../event';
import SearchList from '@/lib/components/search-list';

type Props = {
	previousEvents: Event[];
	upcomingEvents: Event[];
};

const filterEventsBySearch = (events: Event[], query: string) => query
	? events.filter(e => e.title.toLowerCase().includes(query.toLowerCase()))
	: events

const EventList = ({ previousEvents, upcomingEvents }: Props) => {
	let [search, setSearch] = useState("");

	let filteredEvents = useMemo(() =>
		filterEventsBySearch([...previousEvents, ...upcomingEvents], search),
		[search, previousEvents, upcomingEvents]
	)

	return (
		<SearchList search={search} onChange={(newSearch) => setSearch(newSearch)}>
			{search
				? filteredEvents.map(event => <EventEntry key={event.id} event={event} />)
				: upcomingEvents.map(event => <EventEntry key={event.id} event={event} />)}

			<details>
				<summary className="mb-4">Tidligere events ({previousEvents.length})</summary>
				{previousEvents.map(event => <EventEntry key={event.id} event={event} />)}
			</details>
		</SearchList >
	)
}

export default EventList
