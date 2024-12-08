import type { Readable } from "stream"

export interface ObjectStorage {
  /**
  * @description  Uploads an object to the storage, returning the path to where the object
  * can be accessed
  *
  * @param {string} key - Where the object should be stored, including its name.
  * @param {Buffer} data - The data to be stored in the object.
  */
  uploadObject(key: string, data: Buffer): Promise<string>

  /**
  * @description Gets an object from the object storage.
  * @param {string} key - The key of where to find the object.
  */
  getObject(key: string): Promise<ReadableStream>

  /**
  * @description Deletes an object from the object storage.
  * @param {string} key - The key of where to find the object to be deleted.
  */
  deleteObject(key: string): Promise<void>
}
