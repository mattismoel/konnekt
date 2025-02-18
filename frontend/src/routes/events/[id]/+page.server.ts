import { BACKEND_URL } from "$env/static/private";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { eventSchema } from "$lib/event";

export const load: PageServerLoad = async ({ params }) => {
  const id = parseInt(params.id)

  const res = await fetch(`${BACKEND_URL}/events/${id}`)

  if (!res.ok) {
    return error(500, "Could not load event")
  }

  const event = eventSchema.parse(await res.json())

  return {
    event
  }
}
