import type { CreateVenueDTO, VenueDTO } from "@/dto/venue.dto"
import { type VenueRepository } from "./venue.repository"
import { db } from "@/shared/db/db"
import { getVenueByIDTx, insertVenueTx } from "@/shared/db/venue"

export const createSQLiteVenueRepository = (): VenueRepository => {
  const insertVenue = async (data: CreateVenueDTO): Promise<VenueDTO> => {
    return await db.transaction(async (tx) => {
      const insertedVenue = await insertVenueTx(tx, data)
      return insertedVenue
    })
  }

  const getVenueByID = async (id: number): Promise<VenueDTO | null> => {
    return await db.transaction(async (tx) => {
      const venue = await getVenueByIDTx(tx, id)
      return venue
    })
  }

  return {
    insertVenue,
    getVenueByID,
  }
}
