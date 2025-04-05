import type { PageServerLoad } from "./$types";
import { error, redirect } from "@sveltejs/kit";
import { APIError } from "$lib/error";
import { eventById } from "$lib/event";
import { listArtists } from "$lib/artist";
import { listVenues } from "$lib/venue";
import { hasPermissions } from "$lib/auth";

export const load: PageServerLoad = async ({ locals, url, request }) => {
  if (!hasPermissions(locals.permissions, ["view:event", "edit:event"])) {
    return redirect(302, "/auth/login")
  }

  try {
    const artistsResult = await listArtists()
    const venuesResult = await listVenues({
      credentials: "include",
      headers: request.headers,
    })

    const id = url.searchParams.get("id")
    if (!id) {
      return {
        event: null,
        venues: venuesResult.records,
        artists: artistsResult.records,
      }
    }

    const event = await eventById(parseInt(id))

    return {
      event,
      venues: venuesResult.records,
      artists: artistsResult.records,
    }
  } catch (e) {
    if (e instanceof APIError) {
      return error(e.status, e.message)
    }

    return error(500, "Could not load")
  }
}
