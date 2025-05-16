import Caroussel from '@/lib/components/caroussel';
import EventCard from '@/lib/components/event-card';
import EventDetails from '@/lib/components/event-details';
import { useListUpcomingEvents } from '@/lib/features/hook';

import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/events/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { data, isLoading } = useListUpcomingEvents()
  if (isLoading) <p>Loading...</p>
  return (
    <main className="min-h-svh">
      {(data?.records && data.records.length > 0) && (
        <EventDetails event={data.records[0]} prefix="Næste event:" />
      )}
      <div className="px-auto flex flex-col gap-16 pt-32 pb-16">
        <section className="flex flex-col">
          <h1 className="font-heading mb-4 text-5xl font-bold md:text-7xl">Events</h1>
          <span className="text-text/75"> Her kan du se alle kommende events. </span>
        </section>
        {(!data?.records || data.records.length <= 0) && (
          <span className="text-text/50 italic">Der er ingen kommende events i øjeblikket...</span>
        )}
        <Caroussel>
          {data?.records.map(event => (
            <EventCard key={event.id} event={event} />
          ))}
        </Caroussel>
      </div>
    </main>
  )
}
