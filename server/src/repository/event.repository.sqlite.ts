import type { CreateEventDTO, EventDTO } from "@/dto/event.dto";
import { type EventRepository } from "./event.repository";
import { db } from "@/shared/db/db";
import { deleteEventTx, getEventByIDTx, insertEventTx, listEventsTx, setEventCoverImageUrlTx } from "@/shared/db/event";
import type { EventListResult, EventQueryOpts } from "@/shared/event/event";

export const createSQLiteEventRepository = (): EventRepository => {
  const insertEvent = async (data: CreateEventDTO): Promise<EventDTO> => {
    return await db.transaction(async (tx) => {
      return await insertEventTx(tx, data)
    })
  }

  const deleteEvent = async (id: number): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await deleteEventTx(tx, id)
    })
  }

  const getEventByID = async (id: number): Promise<EventDTO | null> => {
    return await db.transaction(async (tx) => {
      return await getEventByIDTx(tx, id)
    })
  }

  const setCoverImageUrl = async (eventID: number, coverImageUrl: string): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await setEventCoverImageUrlTx(tx, eventID, coverImageUrl)
    })
  }

  const listEvents = async (opts: EventQueryOpts): Promise<EventListResult> => {
    return await db.transaction(async (tx) => {
      return await listEventsTx(tx, opts)
    })
  }

  return {
    insertEvent,
    getEventByID,
    listEvents,
    deleteEvent,
    setCoverImageUrl
  }
}
