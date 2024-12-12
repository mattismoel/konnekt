import type { ObjectStorage } from "@/shared/object-storage/object-storage";
import { createEventSchema, type CreateEventDTO, type EventDTO } from "@/dto/event.dto";
import { type EventRepository } from "@/repository/event.repository";
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

  uploadCoverImage = async (buffer: Buffer): Promise<string> => {
    const key = `events/${crypto.randomUUID()}.jpeg`

    const image = await sharp(buffer)
      .resize({ fit: "cover", width: 2048 })
      .jpeg()
      .toBuffer()

    await this.objectStorage.deleteObject(key)
    const coverImageUrl = await this.objectStorage.uploadObject(key, image)

    return coverImageUrl
  }
}
