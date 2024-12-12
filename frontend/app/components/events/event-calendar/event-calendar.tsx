import { addDays, differenceInDays, set } from "date-fns";
import { EventDTO } from "@/lib/dto/event.dto";
import { getEarliestEvent, getLatestEvent } from "@/lib/event";
import { formatDateString, formatHoursAsTimestamp, getFullHoursSurroudningDates } from "@/lib/time";
import { cn } from "@/lib/utils";
import { EventColumn } from "./event-column";

type Props = {
  events?: EventDTO[]
  activeID?: number;
  className?: string;
}

const getEventsOfDate = (date: Date, events: EventDTO[]): EventDTO[] => {
  return events.filter(event => (
    event.fromDate.getDate() === date.getDate()
  ));
}

export const EventCalendar = ({ events = [], activeID, className }: Props) => {
  const firstEvent = events.at(0)
  const earliestEvent = getEarliestEvent(events)
  const latestEvent = getLatestEvent(events)

  const daysToDisplay = (latestEvent && earliestEvent)
    ? differenceInDays(latestEvent.toDate, earliestEvent.fromDate) + 1
    : undefined;

  const hoursToDisplay = (latestEvent && earliestEvent)
    ? getFullHoursSurroudningDates(earliestEvent.fromDate, latestEvent.toDate)
    : undefined;

  if (!earliestEvent || !latestEvent) return (
    <div
      className={cn(`w-full h-full flex items-center justify-center italic text-neutral-500`, className)}
    >
      Ingen øvrige events i denne uge...
    </div>
  )

  return (
    <div
      className={cn("relative grid gap-x-8 w-full overflow-scroll rounded-md", className)}
      style={{
        gridTemplateColumns: `64px repeat(${daysToDisplay}, minmax(256px, 1fr))`,
        gridTemplateRows: `4rem repeat(${hoursToDisplay}, 1fr)`
      }}
    >
      {[...Array(daysToDisplay)].map((_, i) => {
        const date = addDays(firstEvent?.fromDate || new Date(), i)

        return (
          <div className="font-bold text-xl text-neutral-500"
            style={{ gridColumn: 2 + i }}
          >
            {formatDateString(date)}
          </div>
        )
      })}
      {[...Array(hoursToDisplay)].map((_, i) =>
        <div
          className="pt-2 border-t border-neutral-500 text-neutral-500"
          style={{ gridColumn: 1, gridRow: 2 + i }}
        >
          {formatHoursAsTimestamp(earliestEvent.fromDate.getHours() + i)}
        </div>
      )}
      {[...Array(daysToDisplay)].map((_, i) => (
        <div
          style={{
            gridColumn: 2 + i,
            gridRow: `${2} / span ${hoursToDisplay}`
          }}
        >
          <EventColumn
            activeID={activeID}
            startDate={set(earliestEvent.fromDate || new Date(), { minutes: 0 })}
            hourCount={hoursToDisplay || 0}
            events={firstEvent
              ? getEventsOfDate(addDays(firstEvent.fromDate, i), events)
              : []}
          />
        </div>
      ))}
    </div >
  )
}
