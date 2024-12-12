import type { ObjectStorage } from "@/shared/object-storage/object-storage";
import { createEventSchema, type CreateEventDTO, type EventDTO } from "@/dto/event.dto";
import { type EventRepository } from "@/repository/event.repository";
import sharp from "sharp";
import type { EventQueryOpts } from "@/shared/event/event";

export type EventService = {
  createEvent(eventData: CreateEventDTO): Promise<EventDTO>
  deleteEvent(id: number): Promise<void>
  listEvents(opts: EventQueryOpts): Promise<EventDTO[]>
  getEventByID(id: number): Promise<EventDTO | null>
  uploadCoverImage(buffer: Buffer): Promise<string>
}

export const createEventService = (
  eventRepo: EventRepository,
  objectStorage: ObjectStorage,
): EventService => {
  const createEvent = async (eventData: CreateEventDTO): Promise<EventDTO> => {
    const data = createEventSchema.parse(eventData)

    const event = await eventRepo.insertEvent(data)

    return event
  }

  const deleteEvent = async (id: number): Promise<void> => {
    await eventRepo.deleteEvent(id)
  }

  const listEvents = async (opts: EventQueryOpts): Promise<EventDTO[]> => {
    const events = await eventRepo.listEvents(opts)
    return events
  }

  const getEventByID = async (id: number): Promise<EventDTO | null> => {
    const event = await eventRepo.getEventByID(id)

    return event
  }

  const uploadCoverImage = async (buffer: Buffer): Promise<string> => {
    const key = `events/${crypto.randomUUID()}.jpeg`

    const image = await sharp(buffer)
      .resize({ fit: "cover", width: 2048 })
      .jpeg()
      .toBuffer()

    await objectStorage.deleteObject(key)
    const coverImageUrl = await objectStorage.uploadObject(key, image)

    return coverImageUrl
  }

  return {
    createEvent,
    deleteEvent,
    listEvents,
    getEventByID,
    uploadCoverImage
  }
}
