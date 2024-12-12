import type { CreateVenueDTO, VenueDTO } from "@/dto/venue.dto"

export interface VenueRepository {
  insertVenue(data: CreateVenueDTO): Promise<VenueDTO>
  getVenueByID(id: number): Promise<VenueDTO | null>
}
