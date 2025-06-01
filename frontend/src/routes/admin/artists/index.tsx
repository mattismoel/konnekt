import AdminHeader from '@/lib/components/admin-header'
import LinkButton from '@/lib/components/ui/button/link-button'
import { useAuth } from '@/lib/context/auth'
import ArtistList from '@/lib/features/artist/components/artist-list'
import { artistsQueryOpts, upcomingArtistsQueryOpts } from '@/lib/features/artist/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { FaPlus } from 'react-icons/fa6'

export const Route = createFileRoute('/admin/artists/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(upcomingArtistsQueryOpts)
    queryClient.ensureQueryData(artistsQueryOpts)
  }
})

function RouteComponent() {
  const { data: upcomingArtists } = useSuspenseQuery(upcomingArtistsQueryOpts)
  const { data: { records: artists } } = useSuspenseQuery(artistsQueryOpts)

  const { hasPermissions } = useAuth()

  return (
    <>
      <AdminHeader>
        <AdminHeader.Title>Kunstnere</AdminHeader.Title>
        <AdminHeader.Description
        >Overblik over alle kunstnere, som er associerede med events.</AdminHeader.Description
        >
        <AdminHeader.Actions>
          {hasPermissions(["edit:artist"]) && (
            <LinkButton
              to="/admin/artists/create"
              disabled={!hasPermissions(['edit:artist'])}
            >
              <FaPlus />Tilføj
            </LinkButton>
          )}
        </AdminHeader.Actions>
      </AdminHeader>

      <main className="pt-16">
        {hasPermissions(['view:artist']) ? (
          <section className="space-y-4">
            <ArtistList
              artists={artists}
              upcomingArtists={upcomingArtists}
            />
          </section>
        ) : (
          <span>Du har ikke tilladelse til at se kunstnere...</span>

        )}
      </main>
    </>
  )
}
