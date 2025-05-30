import PageMeta from '@/lib/components/page-meta';
import EventDetails from '@/lib/features/event/components/event-details';
import EventGrid from '@/lib/features/event/components/event-grid';
import { upcomingEventsQueryOpts } from '@/lib/features/event/query';
import { useSuspenseQuery } from '@tanstack/react-query';

import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/events/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(upcomingEventsQueryOpts())
  },
})

function RouteComponent() {
  const { data: { records: events } } = useSuspenseQuery(upcomingEventsQueryOpts())

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
        <div className="px-auto flex flex-col pt-8 md:pt-16 pb-16">
          <h1 className="font-heading mb-8 font-bold text-4xl">Kommende events</h1>
          {(events.length <= 0) && (
            <span className="text-text/50 italic">Der er ingen kommende events i øjeblikket...</span>
          )}
          <EventGrid events={events} />
        </div>
      </main>
    </>
  )
}
