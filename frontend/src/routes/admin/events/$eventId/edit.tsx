import { APIError } from '@/lib/api'
import EventForm from '@/lib/components/event-form/event-form'
import { useToast } from '@/lib/context/toast'
import { updateEvent, type editEventForm } from '@/lib/features/event'
import { useArtists, useEventById, useVenues } from '@/lib/features/hook'
import { useQueryClient } from '@tanstack/react-query'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import type { z } from 'zod'

export const Route = createFileRoute('/admin/events/$eventId/edit')({
  component: RouteComponent,
})

function RouteComponent() {
  const { eventId } = Route.useParams()
  const { addToast } = useToast()
  const queryClient = useQueryClient()
  const navigate = useNavigate()

  const eventQuery = useEventById(parseInt(eventId))
  const artistsQuery = useArtists()
  const venuesQuery = useVenues()

  const onSubmit = async (form: z.infer<typeof editEventForm>) => {
    try {
      await updateEvent(form, parseInt(eventId))
      addToast("Event opdateret")
      await queryClient.invalidateQueries({ queryKey: ["events", parseInt(eventId)] })
      navigate({ to: "/admin/events" })

    } catch (e) {
      if (e instanceof APIError) {
        addToast("Kunne ikke opdatere event", e.message, "error")
        return
      }

      addToast("Kunne ikke opdatere event", "Noget gik galt...", "error")
      throw e
    }
  }

  const isLoading = eventQuery.isLoading || artistsQuery.isLoading || venuesQuery.isLoading
  const isError = eventQuery.isError || artistsQuery.isError || venuesQuery.isError

  if (isLoading) return <p>Loading...</p>
  if (isError) return <p>Error...</p>

  return (
    <EventForm
      event={eventQuery.data}
      venues={venuesQuery.data?.records || []}
      artists={artistsQuery.data?.records || []}
      onSubmit={onSubmit}
    />
  )
}
