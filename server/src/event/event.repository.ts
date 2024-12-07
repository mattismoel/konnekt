import type { CreateEventDTO, EventDTO } from "./event.dto";

export class NotFoundError extends Error {
  constructor(msg: string) {
    super(msg)
    Object.setPrototypeOf(this, NotFoundError.prototype)
  }
}

export interface EventRepository {
  insert(event: CreateEventDTO): Promise<EventDTO>;
  delete(id: number): Promise<void>;
  getByID(id: number): Promise<EventDTO | null>
  getAll(): Promise<EventDTO[]>
}
