import { EventDTO } from "@/lib/dto/event.dto"
import { cn } from "@/lib/utils";
import { EventCard } from "./event-card";

const LOADING_PLACEHOLDER_COUNT = 4

type Props = {
  events: EventDTO[]
  isLoading: boolean;
  className?: string;
}

export const EventGrid = ({ events, isLoading, className }: Props) => {
  if (isLoading) return (
    <div className={cn("grid grid-cols-1 md:grid-cols-2 gap-4", className)}>
      {[...Array(LOADING_PLACEHOLDER_COUNT)].map((_, i) =>
        <EventCard key={i} event={undefined} isLoading={isLoading} />
      )}
    </div>
  )

  return (
    <div className={cn("grid grid-cols-1 md:grid-cols-2 gap-4", className)}>
      {events.map(event => <EventCard key={event.id} event={event} isLoading={false} />)}
    </div>
  )
}
