import { BiCalendar, BiMap, BiMusic } from "react-icons/bi"
import { EventDTO } from "@/lib/dto/event.dto"
import { formatDateString } from "@/lib/time"

type Props = {
  event: EventDTO
}

export const EventDetails = ({ event }: Props) => {
  const { venue, fromDate } = event

  return (
    <div className="space-y-2">
      <div className="flex gap-2 items-center">
        <BiMap />
        <span>{venue.name}, {venue.city}</span>
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
