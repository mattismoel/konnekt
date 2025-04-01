import { listVenues } from "$lib/venue";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ request }) => {
  const { records } = await listVenues({
    credentials: "include",
    headers: request.headers,
  })

  return {
    venues: records
  }
}
