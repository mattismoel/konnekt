import type { CreateEventDTO, EventDTO } from "@/dto/event.dto"
import type { TX } from "./db"
import { eventsGenresTable, eventsTable } from "./schema/event"
import { getEventGenresTx, insertGenresTx, relateGenresToEventTx } from "./genre"
import { eq } from "drizzle-orm"
import { getVenueByIDTx, insertVenueTx } from "./venue"

export const insertEventTx = async (tx: TX, data: CreateEventDTO): Promise<EventDTO> => {
  return await tx.transaction(async (tx) => {
    const venue = await insertVenueTx(tx, data.venue)

    const results = await tx
      .insert(eventsTable)
      .values({
        ...data,
        venueID: venue.id
      })
      .returning()

    const { venueID, ...event } = results[0]

    await insertGenresTx(tx, data.genres)
    await relateGenresToEventTx(tx, event.id, data.genres)

    const genres = await getEventGenresTx(tx, event.id)

    return {
      ...event,
      venue: { ...venue },
      genres: genres.map(genre => genre.name)
    }
  })
}

export const deleteEventTx = async (tx: TX, id: number): Promise<void> => {
  await tx.transaction(async (tx) => {
    const event = await getEventByIDTx(tx, id)

    if (!event) return

    await tx.delete(eventsGenresTable).where(eq(eventsGenresTable.eventID, id))
    await tx.delete(eventsTable).where(eq(eventsTable.id, id))
  })
}

export const getEventByIDTx = async (tx: TX, id: number): Promise<EventDTO | null> => {
  return await tx.transaction(async (tx) => {
    const result = await tx
      .select()
      .from(eventsTable)
      .where(eq(eventsTable.id, id))

    if (result.length <= 0) {
      return null
    }

    const { venueID, ...baseEvent } = result[0]

    const genres = await getEventGenresTx(tx, baseEvent.id)
    const venue = await getVenueByIDTx(tx, venueID)

    if (!venue) {
      throw new Error(`No venue found for event with id ${id}`)
    }

    return {
      ...baseEvent,
      venue,
      genres: genres.map(genre => genre.name)
    }
  })
}

export const setEventCoverImageUrlTx = async (tx: TX, eventID: number, coverImageUrl: string): Promise<void> => {
  await tx
    .update(eventsTable)
    .set({ coverImageUrl })
    .where(eq(eventsTable.id, eventID))
}

export const getAllEventsTx = async (tx: TX): Promise<EventDTO[]> => {
  return await tx.transaction(async (tx) => {
    const baseEvents = await tx
      .select()
      .from(eventsTable)

    const events = await Promise.all(
      baseEvents.map(async (baseEvent): Promise<EventDTO> => {
        const venue = await getVenueByIDTx(tx, baseEvent.venueID)

        if (!venue) {
          throw new Error(`No venue found for event with id ${baseEvent.id}`)
        }

        const genres = await getEventGenresTx(tx, baseEvent.id)

        return {
          ...baseEvent,
          venue,
          genres: genres.map(genre => genre.name)
        }
      }))

    return events
  })
}
