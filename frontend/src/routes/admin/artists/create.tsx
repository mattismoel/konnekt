import { genresQueryOpts } from '@/lib/features/artist/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import ArtistForm from '@/lib/features/artist/components/artist-form'

export const Route = createFileRoute('/admin/artists/create')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(genresQueryOpts)
  }
})

function RouteComponent() {
  const { data: { records: genres } } = useSuspenseQuery(genresQueryOpts)

  return (
    <ArtistForm genres={genres} />
  )
}
