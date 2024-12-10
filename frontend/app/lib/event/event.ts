import env from "~/config/env"
import { EventDTO, eventSchema } from "./event.dto"

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
