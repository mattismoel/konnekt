import { APIError } from '@/lib/api'
import ArtistForm from '@/lib/components/artist-form'
import { useToast } from '@/lib/context/toast'
import { updateArtist, type editArtistForm } from '@/lib/features/artist'
import { useArtistById, useGenres } from '@/lib/features/hook'
import { useQueryClient } from '@tanstack/react-query'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import type { z } from 'zod'

export const Route = createFileRoute('/admin/artists/$artistId/edit')({
  component: RouteComponent,
})

function RouteComponent() {
  const { artistId } = Route.useParams()
  const { addToast } = useToast()
  const queryClient = useQueryClient()

  const navigate = useNavigate()

  const artistQuery = useArtistById(parseInt(artistId))
  const genreQuery = useGenres()

  const isLoading = artistQuery.isLoading || genreQuery.isLoading
  const isError = artistQuery.isError || genreQuery.isError

  if (isLoading) return <p>Loading...</p>
  if (isError) return <p>Error...</p>

  const onSubmit = async (form: z.infer<typeof editArtistForm>) => {
    try {
      await updateArtist(parseInt(artistId), form)
      addToast("Kunstner opdateret")
      await queryClient.invalidateQueries({ queryKey: ["artists"] })
      navigate({ to: "/admin/artists" })
    } catch (e) {
      if (e instanceof APIError) {
        addToast("Kunne ikke opdatere kunstner", e.message, "error")
        throw e
      }

      addToast("Kunne ikke opdatere kunstner", "Noget gik galt...", "error")
      throw e
    }
  }

  return (
    <ArtistForm onSubmit={onSubmit} artist={artistQuery.data} genres={genreQuery.data?.records || []} />
  )
}
