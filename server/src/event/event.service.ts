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
}
