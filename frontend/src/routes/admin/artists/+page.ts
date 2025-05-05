import { eventsArtists, listArtists } from "$lib/features/artist/artist";
import { redirect } from "@sveltejs/kit";
import type { PageLoad } from "./$types";
import { memberSession } from "$lib/features/auth/member";
import { hasSomeTeam } from "$lib/features/auth/team";
import { listUpcomingEvents } from "$lib/features/event/event";

export const load: PageLoad = async ({ fetch }) => {
  const member = await memberSession(fetch)

  const { records: upcomingEvents } = await listUpcomingEvents(fetch)

  const upcomingArtists = eventsArtists(upcomingEvents)

  if (!hasSomeTeam(member.teams, ["admin", "event-management"])) {
    return redirect(302, "/auth/login")
  }

  const { records: artists } = await listArtists(fetch)

  return {
    artists,
    upcomingArtists,
  }
}
