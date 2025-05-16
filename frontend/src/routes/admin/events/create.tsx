import { APIError } from '@/lib/api'
import EventForm from '@/lib/components/event-form/event-form'
import { useToast } from '@/lib/context/toast'
import { createEvent, createEventForm } from '@/lib/features/event'
import { useArtists, useVenues } from '@/lib/features/hook'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import type { z } from 'zod'

export const Route = createFileRoute('/admin/events/create')({
  component: RouteComponent,
})

function RouteComponent() {
  const { addToast } = useToast()
  const navigate = useNavigate()

  const artistQuery = useArtists()
  const venueQuery = useVenues()

  const isLoading = artistQuery.isLoading || venueQuery.isLoading
  const isError = artistQuery.isError || venueQuery.isError

  if (isLoading) return <p>Loading...</p>
  if (isError) return <p>Error...</p>

  const onSubmit = async (form: z.infer<typeof createEventForm>) => {
    try {
      await createEvent(form)
      addToast("Event tilføjet")
      navigate({ to: "/admin/events" })
    } catch (e) {
      console.error(e)
      if (e instanceof APIError) {
        addToast("Noget gik galt...", e.message, "error")
      }
    }
  }

  return (
    <EventForm
      venues={venueQuery.data?.records || []}
      artists={artistQuery.data?.records || []}
      disabled={false}
      onSubmit={onSubmit}
    />
  )
}
