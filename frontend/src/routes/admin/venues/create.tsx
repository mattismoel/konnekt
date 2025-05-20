import { createFileRoute } from '@tanstack/react-router'

import VenueForm from '@/lib/features/event/components/venue-form'

export const Route = createFileRoute('/admin/venues/create')({
  component: RouteComponent,
})

function RouteComponent() {

  return (
    <main>
      <h1 className="font-heading text-2xl font-bold mb-4">Lav venue</h1>
      <VenueForm />
    </main >
  )
}
