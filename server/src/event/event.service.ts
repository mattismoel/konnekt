import { createEventSchema, type CreateEventDTO, type EventDTO } from "./event.dto";
import type { EventRepository } from "./event.repository";


export class EventService {
  constructor(
    private readonly eventRepository: EventRepository
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

  getByID = async (id: number): Promise<EventDTO> => {
    const event = await this.eventRepository.getByID(id)

    if (!event) {
      throw Error(`No event found with id ${id}`)
    }

    return event
  }
}
