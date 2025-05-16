import { APIError } from '@/lib/api'
import ArtistForm from '@/lib/components/artist-form'
import { useToast } from '@/lib/context/toast'
import { createArtist, type createArtistForm } from '@/lib/features/artist'
import { useGenres } from '@/lib/features/hook'
import { useQueryClient } from '@tanstack/react-query'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import type { z } from 'zod'

export const Route = createFileRoute('/admin/artists/create')({
  component: RouteComponent,
})

function RouteComponent() {
  const { addToast } = useToast()
  const navigate = useNavigate()
  const queryClient = useQueryClient()

  const genreQuery = useGenres()

  const genres = genreQuery.data?.records || []

  const onSubmit = async (form: z.infer<typeof createArtistForm>) => {
    try {
      await createArtist(form)
      addToast("Kunstner lavet")
      await queryClient.invalidateQueries({ queryKey: ["artists"] })
      navigate({ to: "/admin/artists" })
    } catch (e) {
      if (e instanceof APIError) {
        addToast("Kunne ikke lave kunstner", e.message, "error")
        throw e
      }

      addToast("Kunne ikke lave kunstner", "Noget gik galt...", "error")
      throw e
    }
  }

  return (
    <ArtistForm genres={genres} onSubmit={onSubmit} />
  )
}
