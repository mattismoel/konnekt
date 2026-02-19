import PageMeta from "@/lib/components/page-meta";
import EventDetails from "@/lib/features/event/components/event-details";
import EventGrid from "@/lib/features/event/components/event-grid";
import { upcomingEventsQueryOpts } from "@/lib/features/event/query";
import { useSuspenseQuery } from "@tanstack/react-query";

import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_app/events/")({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(upcomingEventsQueryOpts());
  },
});

function RouteComponent() {
  const {
    data: { records: events },
  } = useSuspenseQuery(upcomingEventsQueryOpts());

  const eventNames = events.map((e) => e.title);

  return (
    <>
      <PageMeta
        title="Konnekt | Events"
        description={`Her kan du se Konnekts kommende events. Oplev blandt andet events som ${eventNames.join(", ")}`}
      />

      {events.length > 0 ? (
        <main className="min-h-svh">
          <EventDetails event={events[0]} prefix="Næste event:" />

          <div className="border-t border-t-zinc-900">
            <div className="mx-responsive flex flex-col pt-16 pb-16 md:pt-16">
              <h1 className="mb-8 font-heading text-4xl font-bold text-heading">
                Kommende events
              </h1>
              <EventGrid events={events} />
            </div>
          </div>
        </main>
      ) : (
        <main className="mx-responsive flex min-h-svh flex-col items-center justify-center">
          <span className="text-text/75 text-center italic">
            Der er ingen aktuelle events i øjeblikket...
          </span>
        </main>
      )}
    </>
  );
}
