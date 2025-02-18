import { error } from "@sveltejs/kit"
import type { PageServerLoad } from "./$types"
import { createListResult } from "$lib/list-result"
import { eventSchema } from "$lib/event"
import { PUBLIC_BACKEND_URL } from "$env/static/public"

export const load: PageServerLoad = async () => {
  const res = await fetch(`${PUBLIC_BACKEND_URL}/events`)
  if (!res.ok) {
    return error(500, "Could not load events")
  }

  const result = createListResult(eventSchema).parse(await res.json())

  return {
    events: result.records
  }
}

