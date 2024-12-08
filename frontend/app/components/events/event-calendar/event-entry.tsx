import { EventDTO } from "~/lib/event/event.dto";
import { getEventDuration } from "~/lib/event/util";
import { distanceBetweenTimesOfDay } from "~/lib/time/format";

type Props = {
  active: boolean;
  event: EventDTO;
  startDate: Date;
  totalMillis: number;
};

export const EventEntry = ({ event, active, startDate, totalMillis }: Props) => {
  const { id, title, fromDate, toDate } = event;

  let duration = getEventDuration(event);
  let heightRatio = duration / totalMillis;

  let topOffsetRatio = distanceBetweenTimesOfDay(startDate, fromDate) / totalMillis

  let colorMap = new Map<string, string>([
    ["active", "text-green-200 bg-green-950 border-green-700 hover:bg-green-900"],
    ["inactive", "text-blue-400 border-blue-900 hover:bg-blue-950"],
  ]);

  return (
    <a
      href={`/events/${id}`}
      className={`absolute w-full p-2 text-sm rounded-md border 
      transition-colors ${colorMap.get(active ? "active" : "inactive")}`}
      style={{
        height: `calc(${heightRatio} * 100%)`,
        top: `calc(${topOffsetRatio} * 100%)`,
      }}
    >
      {title}, {fromDate.toLocaleTimeString()} - {toDate.toLocaleTimeString()}
    </a>
  )
}
