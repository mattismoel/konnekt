import type { EventDTO } from "@/dto/event.dto"

export type EventQueryOpts = {
  page: number | undefined,
  limit: number | undefined,
  pageSize: number | undefined,
  search: string | undefined,
  fromDate: Date | undefined,
  toDate: Date | undefined,
}

export type EventListResult = {
  totalSize: number;
  events: EventDTO[],
}
