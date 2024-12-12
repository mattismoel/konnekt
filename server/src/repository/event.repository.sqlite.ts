import type { CreateEventDTO, EventDTO } from "@/dto/event.dto";
import { type EventRepository } from "./event.repository";
import { db } from "@/shared/db/db";
import { deleteEventTx, getAllEventsTx, getEventByIDTx, insertEventTx, setEventCoverImageUrlTx } from "@/shared/db/event";

export class SQLiteEventRepository implements EventRepository {
  async insert(data: CreateEventDTO): Promise<EventDTO> {
    return await db.transaction(async (tx) => {
      return await insertEventTx(tx, data)
    })
  }

  async delete(id: number): Promise<void> {
    return await db.transaction(async (tx) => {
      return await deleteEventTx(tx, id)
    })
  }

  getByID = async (id: number): Promise<EventDTO | null> => {
    return await db.transaction(async (tx) => {
      return await getEventByIDTx(tx, id)
    })
  }

  setCoverImageUrl = async (eventID: number, coverImageUrl: string): Promise<void> => {
    return await db.transaction(async (tx) => {
      return await setEventCoverImageUrlTx(tx, eventID, coverImageUrl)
    })
  }

  getAll = async (): Promise<EventDTO[]> => {
    return await db.transaction(async (tx) => {
      return await getAllEventsTx(tx)
    })
  }
}
