import type { CreateVenueDTO, VenueDTO } from "@/dto/venue.dto";
import type { TX } from "./db";
import { venuesTable } from "./schema/venue";
import { eq } from "drizzle-orm";

/**
 * @description Inserts a venue into the database.
 */
export const insertVenueTx = async (tx: TX, venue: CreateVenueDTO): Promise<VenueDTO> => {
  const result = await tx
    .insert(venuesTable)
    .values(venue)
    .returning()

  return result[0]
}

/**
 * @description Returns a venue given its ID. If not found, null is returned.
 */
export const getVenueByIDTx = async (tx: TX, id: number): Promise<VenueDTO | null> => {
  const result = await tx
    .select()
    .from(venuesTable)
    .where(eq(venuesTable.id, id))

  if (result.length <= 0) {
    return null
  }

  return result[0]
}

