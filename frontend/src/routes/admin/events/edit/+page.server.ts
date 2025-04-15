import type { PageServerLoad } from "./$types";
import { redirect } from "@sveltejs/kit";
import { eventById } from "$lib/features/event/event";
import { listArtists } from "$lib/features/artist/artist";
import { listVenues } from "$lib/features/venue/venue";

import { hasPermissions, userPermissions } from "$lib/features/auth/permission";
import { userSession } from "$lib/features/auth/user";

export const load: PageServerLoad = async ({ url, fetch }) => {
  const user = await userSession(fetch)
  const permissions = await userPermissions(fetch, user.id)

  if (!hasPermissions(permissions, ["view:event", "edit:event"])) {
    return redirect(302, "/auth/login")
  }

  const { records: artists } = await listArtists(fetch)
  const { records: venues } = await listVenues(fetch)

  const id = url.searchParams.get("id")
  if (!id) {
    return { event: null, venues, artists }
  }

  const event = await eventById(fetch, parseInt(id))

  return { event, venues, artists }
}
