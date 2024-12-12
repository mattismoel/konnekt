import { differenceInMilliseconds } from "date-fns";
import env from "@/config/env"
import { EventDTO, eventSchema } from "./dto/event.dto";

/**
 * @description Gets a single event by its id. If not found, null is returned.
 */
export const fetchEventByID = async (id: number, opts: RequestInit): Promise<EventDTO | null> => {
  const res = await fetch(`${env.BACKEND_URL}/events/${id}`, opts)
  if (!res.ok) {
    console.error(`Could not get event with id ${id}: ${res.statusText}`)
    return null
  }

  const event = eventSchema.parse(await res.json())

  return event
}

/**
 * @description Gets all events.
 */
export const fetchAllEvents = async (opts: RequestInit): Promise<EventDTO[]> => {
  const res = await fetch(`${env.BACKEND_URL}/events`, opts)
  if (!res.ok) {
    console.error(`Could not list events: ${res.statusText}`)
    return []
  }

  const events = eventSchema.array().parse(await res.json())

  return events
}


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
