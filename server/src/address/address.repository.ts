import type { AddressDTO, CreateAddressDTO } from "./address.dto";

export interface AddressRespository {
  /**
   * @description Inserts an address into the database.
   *
   * @param {CreateAddressDTO} address - The address to insert.
   */
  insertAddress(address: CreateAddressDTO): Promise<AddressDTO>

  /**
   * @description Finds an address matching the input ID.
   *
   * @param {number} id - The ID of an address to find.
   */
  getAddressByID(id: number): Promise<AddressDTO>
}
