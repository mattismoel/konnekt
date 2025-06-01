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

      {(events.length > 0) ? (
        <main className="min-h-svh">
          <EventDetails event={events[0]} prefix="Næste event:" />

          <div className="px-auto flex flex-col pt-16 md:pt-16 pb-16">
            <h1 className="font-heading mb-8 text-4xl font-bold text-shadow-md/15">Kommende events</h1>
            <EventGrid events={events} />
          </div>
        </main>
      ) : (
        <main className="flex flex-col justify-center items-center min-h-svh px-auto">
          <span className="italic text-text/75 text-center">Der er ingen aktuelle events i øjeblikket...</span>
        </main>
      )}
    </>
  )
}
