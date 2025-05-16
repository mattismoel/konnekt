import AdminHeader from '@/lib/components/admin-header'
import ArtistList from '@/lib/components/artist-list'
import LinkButton from '@/lib/components/ui/button/link-button'
import { useAuth } from '@/lib/context/auth'
import { useArtists, useListUpcomingArtists } from '@/lib/features/hook'
import { createFileRoute } from '@tanstack/react-router'
import { FaPlus } from 'react-icons/fa6'

export const Route = createFileRoute('/admin/artists/')({
  component: RouteComponent,
})

function RouteComponent() {
  const upcomingArtistsQuery = useListUpcomingArtists()
  const artistsQuery = useArtists()

  const { hasPermissions } = useAuth()

  const isLoading = upcomingArtistsQuery.isLoading || artistsQuery.isLoading
  const isError = upcomingArtistsQuery.isError || artistsQuery.isError

  if (isLoading) return <p>Loading...</p>
  if (isError) return <p>Error...</p>

  return (
    <>
      <AdminHeader>
        <AdminHeader.Title>Kunstnere</AdminHeader.Title>
        <AdminHeader.Description
        >Overblik over alle kunstnere, som er associerede med events.</AdminHeader.Description
        >
        <AdminHeader.Actions>
          <LinkButton
            to="/admin/artists/create"
            disabled={!hasPermissions(['edit:artist'])}
          >
            <FaPlus />Tilføj
          </LinkButton>
        </AdminHeader.Actions>
      </AdminHeader>

      <main className="pt-16">
        {hasPermissions(['view:artist']) ? (
          <section className="space-y-4">
            <ArtistList
              artists={artistsQuery.data?.records || []}
              upcomingArtists={upcomingArtistsQuery.data || []}
            />
          </section>
        ) : (
          <span>Du har ikke tilladelse til at se kunstnere...</span>

        )}
      </main>
    </>
  )
}
