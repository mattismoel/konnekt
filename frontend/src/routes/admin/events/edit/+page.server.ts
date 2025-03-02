import type { PageServerLoad } from "./$types";
import { eventById } from "$lib/event";
import { createListResult } from "$lib/list-result";
import { venueSchema } from "$lib/venue";
import { listArtists } from "$lib/artist";

export const load: PageServerLoad = async ({ url, request }) => {
  const artistsResult = await listArtists()
    credentials: "include",
    headers: request.headers,
  })

  const id = url.searchParams.get("id")
  if (!id) {
    return { event: null, venues, artists }
  }

  const event = await eventById(parseInt(id))

  return {
    event,
    artists,
    venues
  }
}
