import { BiCalendar, BiMap, BiMusic } from "react-icons/bi"
import { EventDTO } from "@/lib/dto/event.dto"
import { formatDateString } from "@/lib/time"

type Props = {
  event: EventDTO
}

/* 
 * TODO: Add 'venue' property to event (both server and client).
 * Alternatively add 'customName' to address model.
 */

export const EventDetails = ({ event }: Props) => {
  const { venue, fromDate } = event

  return (
    <div className="space-y-2">
      <div className="flex gap-2 items-center">
        <BiMap />
        <span>{`Posten`}, {venue.city}</span>
      </div>
      <div className="flex gap-2 items-center">
        <BiCalendar />
        <span>{formatDateString(fromDate)}</span>
      </div>
      <div className="flex gap-2 items-center">
        <BiMusic />
        <span>{event.genres.join(", ")}</span>
      </div>
    </div>
  )
}
