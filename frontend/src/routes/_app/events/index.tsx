import Caroussel from '@/lib/components/caroussel';
import PageMeta from '@/lib/components/page-meta';
import EventCard from '@/lib/features/event/components/event-card';
import EventDetails from '@/lib/features/event/components/event-details';
import { upcomingEventsQueryOpts } from '@/lib/features/event/query';
import { useSuspenseQuery } from '@tanstack/react-query';

import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/events/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(upcomingEventsQueryOpts)
  },
  pendingComponent: () => <div>Loading...</div>
})

function RouteComponent() {
  const { data: { records: events } } = useSuspenseQuery(upcomingEventsQueryOpts)

  const eventNames = events.map(e => e.title)

  return (
    <>
      <PageMeta
        title="Konnekt | Events"
        description={`Her kan du se Konnekts kommende events. Oplev blandt andet events som ${eventNames.join(", ")}`}
      />

      <main className="min-h-svh">
        {(events.length > 0) && (
          <EventDetails event={events[0]} prefix="Næste event:" />
        )}
        <div className="px-auto flex flex-col gap-16 pt-32 pb-16">
          <section className="flex flex-col">
            <h1 className="font-heading mb-4 text-5xl font-bold md:text-7xl">Events</h1>
            <span className="text-text/75"> Her kan du se alle kommende events. </span>
          </section>
          {(events.length <= 0) && (
            <span className="text-text/50 italic">Der er ingen kommende events i øjeblikket...</span>
          )}
          <Caroussel>
            {events.map(event => (
              <EventCard key={event.id} event={event} />
            ))}
          </Caroussel>
        </div>
      </main>
    </>
  )
}
