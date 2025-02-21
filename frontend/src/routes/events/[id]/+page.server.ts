import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { eventSchema } from "$lib/event";
import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { createListResult } from "$lib/list-result";

export const load: PageServerLoad = async ({ params }) => {
  const id = parseInt(params.id)

  let res = await fetch(`${PUBLIC_BACKEND_URL}/events/${id}`)

  if (!res.ok) {
    return error(500, "Could not load event")
  }

  const event = eventSchema.parse(await res.json())

  res = await fetch(`${PUBLIC_BACKEND_URL}/events`)
  if (!res.ok) {
    return error(500, "Could not load events")
  }


  const recommendedEventsResult = createListResult(eventSchema).parse(await res.json())

  return {
    event,
    recommendedEvents: recommendedEventsResult.records,
  }
}
