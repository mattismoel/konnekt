import Caroussel from '@/lib/components/caroussel'
import EventCalendar from '@/lib/components/event-calendar'
import EventCard from '@/lib/components/event-card'
import EventDetails from '@/lib/components/event-details'
import { useEventById, useListUpcomingEvents } from '@/lib/features/hook'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/events/$eventId')({
  component: RouteComponent,
})

function RouteComponent() {
  const { eventId } = Route.useParams()
  const eventQuery = useEventById(parseInt(eventId))
  const upcomingEventsQuery = useListUpcomingEvents()

  const isLoading = eventQuery.isLoading || upcomingEventsQuery.isLoading
  const isError = eventQuery.isError || upcomingEventsQuery.isError

  if (isLoading) return <p>Loading...</p>
  if (isError) return <p>Error...</p>

  const event = eventQuery.data
  if (!event) return <p>No such event...</p>

  const upcomingEvents = upcomingEventsQuery.data?.records.filter(e => e.id !== event.id) || []

  return (
    <main className="min-h-sub-nav flex flex-col gap-16 pb-16 text-white">
      <EventDetails active event={event} />
      <article className="px-auto space-y-16 pt-8 pb-16">
        <section className="prose prose-invert max-w-none" dangerouslySetInnerHTML={{ __html: event.description }} />
        <EventCalendar event={event} />
        {upcomingEvents.length > 0 && (
          <section>
            <h1 className="mb-4 text-2xl font-bold">Se også</h1>
            <Caroussel>
              {upcomingEvents.map(event => <EventCard key={event.id} event={event} />)}
            </Caroussel>
          </section>
        )}
      </article>
    </main>
  )
}
