import { addHours, differenceInMilliseconds } from "date-fns";
import { EventDTO } from "~/lib/event/event.dto";
import { EventEntry } from "./event-entry";

type Props = {
  events: EventDTO[];
  activeID?: number;
  startDate: Date;
  hourCount: number;
}

export const EventColumn = ({
  events,
  activeID,
  startDate,
  hourCount
}: Props
) => {
  const endDate = addHours(startDate, hourCount)
  const totalMillis = differenceInMilliseconds(endDate, startDate)

  return (
    <div className="relative w-full h-full">
      {events.map(event => {
        return (
          <EventEntry
            active={event.id === activeID}
            event={event}
            startDate={startDate}
            totalMillis={totalMillis}
          />
        )
      })}
    </div>
  )
}
