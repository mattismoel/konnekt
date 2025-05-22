import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { artistsQueryOpts } from '@/lib/features/artist/query'
import { createEventByIdOpts, venuesQueryOpts } from '@/lib/features/event/query'

import EventForm from '@/lib/features/event/components/event-form/event-form'

export const Route = createFileRoute('/admin/events/$eventId/edit')({
  component: RouteComponent,
  loader: async ({ context: { queryClient }, params: { eventId } }) => {
    const eventQueryOptions = createEventByIdOpts(parseInt(eventId))

    queryClient.ensureQueryData(eventQueryOptions)
    queryClient.ensureQueryData(artistsQueryOpts)
    queryClient.ensureQueryData(venuesQueryOpts)

    return { eventQueryOptions }
  }
})

function RouteComponent() {
  const { eventQueryOptions } = Route.useLoaderData()

  const { data: event } = useSuspenseQuery(eventQueryOptions)
  const { data: { records: artists } } = useSuspenseQuery(artistsQueryOpts)
  const { data: { records: venues } } = useSuspenseQuery(venuesQueryOpts)

  return (
    <EventForm event={event} venues={venues} artists={artists} />
  )
}
