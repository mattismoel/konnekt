import { genresTable } from "@/shared/db/schema/genre";
import type { CreateEventDTO, EventDTO } from "./event.dto";
import { NotFoundError, type EventRepository } from "./event.repository";
import { db, type TX } from "@/shared/db/db";
import { eventsGenresTable, eventsTable } from "@/shared/db/schema/event";
import { addressesTable } from "@/shared/db/schema/address";
import { eq, inArray } from "drizzle-orm";
import type { GenreDTO } from "@/event/genre.dto";
import type { AddressDTO, CreateAddressDTO } from "./address.dto";

export class SQLiteEventRepository implements EventRepository {
  async insert(data: CreateEventDTO): Promise<EventDTO> {
    return await db.transaction(async (tx) => {
      const address = await insertAddress(tx, data.address)

      const results = await tx
        .insert(eventsTable)
        .values({ ...data, addressID: address.id })
        .returning()

      const { addressID, ...event } = results[0]

      await insertGenres(tx, data.genres)
      await relateGenresToEvent(tx, event.id, data.genres)
      const genres = await getEventGenres(tx, event.id)

      return {
        ...event,
        address: { ...address },
        genres: genres.map(genre => genre.name)
      }
    })
  }

  async delete(id: number): Promise<void> {
    await db.transaction(async (tx) => {
      const eventAddress = await getEventAddress(tx, id)

      await tx.delete(eventsGenresTable).where(eq(eventsGenresTable.eventID, id))
      await tx.delete(eventsTable).where(eq(eventsTable.id, id))
      await tx.delete(addressesTable).where(eq(addressesTable.id, eventAddress.id))
    })
  }

  getByID = async (id: number): Promise<EventDTO | null> => {
    return await db.transaction(async (tx) => {
      const result = await tx
        .select()
        .from(eventsTable)
        .where(eq(eventsTable.id, id))

      if (result.length <= 0) {
        throw new NotFoundError(`No events found with id ${id}`)
      }

      const { addressID, ...baseEvent } = result[0]

      const genres = await getEventGenres(tx, baseEvent.id)
      const address = await getEventAddress(tx, baseEvent.id)

      return {
        ...baseEvent,
        address,
        genres: genres.map(genre => genre.name)
      }
    })
  }

  getAll = async (): Promise<EventDTO[]> => {
    return await db.transaction(async (tx) => {
      const baseEvents = await tx
        .select()
        .from(eventsTable)

      const events = await Promise.all(
        baseEvents.map(async (baseEvent): Promise<EventDTO> => {
          const address = await getEventAddress(tx, baseEvent.id)
          const genres = await getEventGenres(tx, baseEvent.id)

          return {
            ...baseEvent,
            address,
            genres: genres.map(genre => genre.name)
          }
        }))
      return events
    })
  }
}

const insertAddress = async (tx: TX, address: CreateAddressDTO): Promise<AddressDTO> => {
  const result = await tx
    .insert(addressesTable)
    .values({ ...address })
    .returning()

  return result[0]
}

const insertGenres = async (tx: TX, genreNames: string[]): Promise<GenreDTO[]> => {
  const genres = await tx
    .insert(genresTable)
    .values(genreNames.map(name => ({ name })))
    .onConflictDoNothing()
    .returning()

  return genres
}

const getEventGenres = async (tx: TX, eventID: number): Promise<GenreDTO[]> => {
  const results = await tx.
    select()
    .from(genresTable)
    .innerJoin(eventsGenresTable, eq(genresTable.id, eventsGenresTable.genreID))
    .where(eq(eventsGenresTable.eventID, eventID))

  const genres = results.map(res => res.genre)

  return genres
}

const relateGenresToEvent = async (tx: TX, eventID: number, genreNames: string[]): Promise<void> => {
  const genres = await tx
    .select({ id: genresTable.id })
    .from(genresTable)
    .where(inArray(genresTable.name, genreNames))

  await tx
    .insert(eventsGenresTable)
    .values(genres.map(genre => ({ eventID, genreID: genre.id })))
}

const getEventAddress = async (tx: TX, id: number): Promise<AddressDTO> => {
  const result = await tx
    .select()
    .from(eventsTable)
    .where(eq(eventsTable.id, id))

  if (result.length <= 0) {
    throw new Error(`Could not find event with id ${id}`)
  }

  const { addressID } = result[0]

  const address = await getAddressByID(tx, addressID)
  return address
}

const getAddressByID = async (tx: TX, id: number): Promise<AddressDTO> => {
  const result = await tx
    .select()
    .from(addressesTable)
    .where(eq(addressesTable.id, id))

  return result[0]
}
