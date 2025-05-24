import type { Event } from "../event";
import EventCard from "./event-card";

type Props = {
	events: Event[]
}

const EventGrid = ({ events }: Props) => (
	<div className="@container">
		<div className="grid gap-4 grid-cols-1 @2xl:grid-cols-2 @5xl:grid-cols-3 @7xl:grid-cols-4">
			{events.map(event => <EventCard key={event.id} event={event} />)}
		</div>
	</div>
)

export default EventGrid
