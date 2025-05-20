import Caroussel from '@/lib/components/caroussel'
import EventCalendar from '@/lib/features/event/components/event-calendar'
import EventCard from '@/lib/features/event/components/event-card'
import EventDetails from '@/lib/features/event/components/event-details'
import { createEventByIdOpts, upcomingEventsQueryOpts } from '@/lib/features/event/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/events/$eventId')({
  component: RouteComponent,
  loader: async ({ context: { queryClient }, params: { eventId: eventIdStr } }) => {
    const eventId = parseInt(eventIdStr)

    const eventOptions = createEventByIdOpts(eventId)

    queryClient.ensureQueryData(eventOptions)
    queryClient.ensureQueryData(upcomingEventsQueryOpts)

    return { eventOptions }
  }
})

function RouteComponent() {
  const { eventOptions } = Route.useLoaderData()

  const { data: event } = useSuspenseQuery(eventOptions)
  const { data: { records: upcomingEvents } } = useSuspenseQuery(upcomingEventsQueryOpts)

  const filteredEvents = upcomingEvents.filter(({ id }) => id !== event.id)

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
              {filteredEvents.map(event => <EventCard key={event.id} event={event} />)}
            </Caroussel>
          </section>
        )}
      </article>
    </main>
  )
}
