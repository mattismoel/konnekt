import { EventDTO } from "@/lib/dto/event.dto"
import { EventCard } from "./event-card"

type Props = {
  events: EventDTO[]
}

export const EventCaroussel = ({ events }: Props) => {
  return (
    <div className="relative flex gap-8 items-center overflow-x-scroll snap-x w-full pb-4">
      {events.map(event => {
        return (
          <div key={event.id} className="scroll-pl-2 h-full sm:w-72 md:w-96 flex-shrink-0 snap-center">
            <EventCard event={event} isLoading={false} />
          </div>
        )
      })}
      <div className="h-64 w-16 flex-shrink-0 snap-center flex justify-center items-center">
        <a href="/events" className="w-full text-center">Se alle.</a>
      </div>
    </div>
  )
}
