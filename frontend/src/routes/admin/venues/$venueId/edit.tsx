import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { createVenueByIdQueryOptions } from '@/lib/features/event/query'

import VenueForm from '@/lib/features/event/components/venue-form'

export const Route = createFileRoute('/admin/venues/$venueId/edit')({
  component: RouteComponent,
  loader: async ({ context: { queryClient }, params: { venueId } }) => {
    const venueQueryOptions = createVenueByIdQueryOptions(parseInt(venueId))

    queryClient.ensureQueryData(venueQueryOptions)

    return { venueQueryOptions }
  }
})

function RouteComponent() {
  const { venueQueryOptions } = Route.useLoaderData()
  const { data: venue } = useSuspenseQuery(venueQueryOptions)

  return (
    <main>
      <h1 className="font-heading text-2xl font-bold mb-4">Redigér venue</h1>
      <VenueForm venue={venue} />
    </main>
  )
}
