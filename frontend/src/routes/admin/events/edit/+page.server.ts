import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { eventById } from "$lib/event";
import { artistSchema } from "$lib/artist";
import { createListResult } from "$lib/list-result";
import { venueSchema } from "$lib/venue";

export const load: PageServerLoad = async ({ url, request }) => {
  let res = await fetch(`${PUBLIC_BACKEND_URL}/artists`, {
    credentials: "include",
    headers: request.headers,
  })

  if (!res.ok) {
    return error(500, "Could not get artists")
  }

  const artists = createListResult(artistSchema).parse(await res.json()).records

  res = await fetch(`${PUBLIC_BACKEND_URL}/venues`, {
    credentials: "include",
    headers: request.headers,
  })

  if (!res.ok) {
    return error(500, "Could not get venues")
  }

  const venues = createListResult(venueSchema).parse(await res.json()).records

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
