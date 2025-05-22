import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { createArtistByIdOpts, genresQueryOpts } from '@/lib/features/artist/query'

import ArtistForm from '@/lib/features/artist/components/artist-form'

export const Route = createFileRoute('/admin/artists/$artistId/edit')({
  component: RouteComponent,
  loader: async ({ context: { queryClient }, params: { artistId } }) => {
    const artistQueryOptions = createArtistByIdOpts(parseInt(artistId))

    queryClient.ensureQueryData(artistQueryOptions)
    queryClient.ensureQueryData(genresQueryOpts)

    return { artistQueryOptions }
  }
})

function RouteComponent() {
  const { artistQueryOptions } = Route.useLoaderData()
  const { data: artist } = useSuspenseQuery(artistQueryOptions)
  const { data: { records: genres } } = useSuspenseQuery(genresQueryOpts)

  return (
    <ArtistForm artist={artist} genres={genres} />
  )
}
