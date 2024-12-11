import { EventDTO } from "@/lib/dto/event.dto"
import { cn } from "@/lib/utils";
import { EventCard } from "./event-card";

type Props = {
  events: EventDTO[]
  className?: string;
}

export const EventGrid = ({ events, className }: Props) => {
  return (
    <div className={cn("grid grid-cols-1 md:grid-cols-2 gap-4", className)}>
      {events.map(event => <EventCard key={event.id} event={event} />)}
    </div>
  )
}
