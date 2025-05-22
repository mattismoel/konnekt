import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { artistsQueryOpts } from '@/lib/features/artist/query'
import { venuesQueryOpts } from '@/lib/features/event/query'

import EventForm from '@/lib/features/event/components/event-form/event-form'

export const Route = createFileRoute('/admin/events/create')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(artistsQueryOpts)
    queryClient.ensureQueryData(venuesQueryOpts)
  }
})

function RouteComponent() {
  const { data: { records: artists } } = useSuspenseQuery(artistsQueryOpts)
  const { data: { records: venues } } = useSuspenseQuery(venuesQueryOpts)

  return (
    <EventForm
      venues={venues}
      artists={artists}
      disabled={false}
    />
  )
}
