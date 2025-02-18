import { PUBLIC_BACKEND_URL } from "$env/static/public";
import { error } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import { eventSchema } from "$lib/event";

export const load: PageServerLoad = async ({ url }) => {
  const id = url.searchParams.get("id")

  if (!id) {
    return {
      event: null
    }
  }

  const res = await fetch(`${PUBLIC_BACKEND_URL}/events/${id}`)
  if (!res.ok) {
    return error(400, "Could not find event")
  }

  const event = eventSchema.parse(await res.json())

  return { event }
}
