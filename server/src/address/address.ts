import { type TX } from "@/shared/db/db"
import { addressesTable } from "@/shared/db/schema/address"
import type { AddressDTO, CreateAddressDTO } from "./address.dto"
import { eq } from "drizzle-orm"

export const insertAddress = async (tx: TX, address: CreateAddressDTO): Promise<AddressDTO> => {
  const result = await tx
    .insert(addressesTable)
    .values({ ...address })
    .returning()

  return result[0]
}

export const getAddressByID = async (tx: TX, id: number): Promise<AddressDTO> => {
  const result = await tx
    .select()
    .from(addressesTable)
    .where(eq(addressesTable.id, id))

  return result[0]
}
