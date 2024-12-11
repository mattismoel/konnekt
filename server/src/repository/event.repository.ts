import type { CreateEventDTO, EventDTO } from "@/dto/event.dto";

export interface EventRepository {
  insert(event: CreateEventDTO): Promise<EventDTO>;
  delete(id: number): Promise<void>;
  getByID(id: number): Promise<EventDTO | null>
  getAll(): Promise<EventDTO[]>
  setCoverImageUrl(eventID: number, url: string): Promise<void>
}
