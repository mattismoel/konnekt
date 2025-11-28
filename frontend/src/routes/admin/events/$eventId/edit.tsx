import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

import { artistsQueryOpts } from "@/lib/features/artist/query";
import {
  createEventByIdOpts,
  venuesQueryOpts,
} from "@/lib/features/event/query";

import EventForm from "@/lib/features/event/components/event-form/event-form";
import { createMemberByIdQueryOpts } from "@/lib/features/auth/query";

export const Route = createFileRoute("/admin/events/$eventId/edit")({
  component: RouteComponent,
  loader: async ({ context: { queryClient }, params: { eventId } }) => {
    const eventQueryOptions = createEventByIdOpts(parseInt(eventId));
    const event = await queryClient.ensureQueryData(eventQueryOptions);

    const updatedByQueryOpts = createMemberByIdQueryOpts(event.updatedBy);

    queryClient.ensureQueryData(updatedByQueryOpts);
    queryClient.ensureQueryData(artistsQueryOpts);
    queryClient.ensureQueryData(venuesQueryOpts);

    return { eventQueryOptions, updatedByQueryOpts };
  },
});

function RouteComponent() {
  const { eventQueryOptions, updatedByQueryOpts } = Route.useLoaderData();

  const { data: event } = useSuspenseQuery(eventQueryOptions);

  const {
    data: { records: artists },
  } = useSuspenseQuery(artistsQueryOpts);

  const {
    data: { records: venues },
  } = useSuspenseQuery(venuesQueryOpts);

  const { data: updatedByMember } = useSuspenseQuery(updatedByQueryOpts);

  return (
    <main className="mx-responsive min-h-svh py-32">
      <EventForm
        event={event}
        venues={venues}
        artists={artists}
        updatedByMember={updatedByMember}
      />
    </main>
  );
}
