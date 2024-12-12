import type { CreateEventDTO, EventDTO } from "@/dto/event.dto";
import type { EventListResult, EventQueryOpts } from "@/shared/event/event";

export interface EventRepository {
  insertEvent(event: CreateEventDTO): Promise<EventDTO>;
  deleteEvent(id: number): Promise<void>;
  getEventByID(id: number): Promise<EventDTO | null>
  listEvents(opts: EventQueryOpts): Promise<EventListResult>
  setCoverImageUrl(eventID: number, url: string): Promise<void>
}
