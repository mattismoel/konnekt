import { differenceInMilliseconds } from "date-fns";
import { EventDTO, EventListResult, eventListSchema, eventSchema } from "./dto/event.dto";
import { sleep } from "./time";

type EventQueryOpts = {
  limit?: number;
  page?: number;
  pageSize?: number;
  search?: string;
}

/**
 * @description Gets a single event by its id. If not found, null is returned.
 */
export const fetchEventByID = async (id: number): Promise<EventDTO | null> => {
  const res = await fetch(`${window.ENV.BACKEND_URL}/events/${id}`, {
    credentials: "include"
  })

  if (!res.ok) {
    console.error(`Could not get event with id ${id}: ${res.statusText}`)
    return null
  }

  const event = eventSchema.parse(await res.json())

  // INFO: ARTIFICIAL DELAY
  //await sleep(1000)

  return event
}

/**
 * @description Gets all events.
 */
export const fetchEvents = async (opts?: EventQueryOpts): Promise<EventListResult> => {
  const { page, limit, pageSize, search } = opts || {}
  const url = new URL(`${window.ENV.BACKEND_URL}/events`)

  if (page) url.searchParams.set("page", page.toString())
  if (limit) url.searchParams.set("limit", limit.toString())
  if (pageSize) url.searchParams.set("pageSize", pageSize.toString())
  if (search) url.searchParams.set("search", search)

  const res = await fetch(url)
  if (!res.ok) {
    console.error(`Could not list events: ${res.statusText}`)
    return { events: [], totalSize: 0 }
  }

  // INFO: Artificial delay
  await sleep(3000)

  const result = eventListSchema.parse(await res.json())

  return result
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
