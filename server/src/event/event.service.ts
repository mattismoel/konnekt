import type { ObjectStorage } from "@/shared/object-storage/object-storage";
import { createEventSchema, type CreateEventDTO, type EventDTO } from "./event.dto";
import { type EventRepository } from "./event.repository";
import sharp from "sharp";


export class EventService {
  constructor(
    private readonly eventRepository: EventRepository,
    private readonly objectStorage: ObjectStorage
  ) { }

  async create(eventData: CreateEventDTO): Promise<EventDTO> {
    const data = createEventSchema.parse(eventData)

    const event = await this.eventRepository.insert(data)

    return event
  }

  async delete(id: number): Promise<void> {
    await this.eventRepository.delete(id)
  }

  getAll = async (): Promise<EventDTO[]> => {
    const events = await this.eventRepository.getAll()
    return events
  }

  getByID = async (id: number): Promise<EventDTO | null> => {
    const event = await this.eventRepository.getByID(id)

    return event
  }

  setCoverImage = async (id: number, buffer: Buffer) => {
    const key = `events/${id}/cover.jpeg`

    const image = await sharp(buffer)
      .resize({ fit: "cover", width: 2048 })
      .jpeg()
      .toBuffer()

    await this.objectStorage.deleteObject(key)
    const coverImageUrl = await this.objectStorage.uploadObject(key, image)

    await this.eventRepository.setCoverImageUrl(id, coverImageUrl)
  }
}
