import { differenceInMilliseconds } from "date-fns";
import { EventDTO } from "./event.dto";

/**
  * @description Returns the total duration of an input event in milliseconds.
  *
  * @param event - The event to return the duration of.
  */
export const getEventDuration = (event: EventDTO): number => {
  return differenceInMilliseconds(event.toDate, event.fromDate);
};

/**
 * Returns the event that is of the earliest time of day.
 *
 * @param events - All events.
 */
export const getEarliestEvent = (events: EventDTO[]): EventDTO | undefined => {
  if (events.length <= 0) return undefined

  return events.reduce((earliest, current) => {
    const earliestTime = earliest.fromDate.getHours() * 60 + earliest.fromDate.getMinutes()
    const currentTime = current.fromDate.getHours() * 60 + current.fromDate.getMinutes()

    return currentTime < earliestTime ? current : earliest
  })
};

/**
 * Returns the event that is of the latest time of day.
 */
export const getLatestEvent = (events: EventDTO[]): EventDTO | undefined => {
  if (events.length <= 0) return undefined

  return events.reduce((a, b) => (a.toDate > b.toDate ? a : b));
};
