import AdminHeader from '@/lib/components/admin-header'
import LinkButton from '@/lib/components/ui/button/link-button'
import { useAuth } from '@/lib/context/auth'
import VenueList from '@/lib/features/event/components/venue-list'
import { venuesQueryOpts } from '@/lib/features/event/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { FaPlus } from 'react-icons/fa'

export const Route = createFileRoute('/admin/venues/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(venuesQueryOpts)
  }
})

function RouteComponent() {
  const { hasPermissions } = useAuth()
  const { data: { records: venues } } = useSuspenseQuery(venuesQueryOpts)
  return (

    <>
      <AdminHeader>
        <AdminHeader.Title>Venues</AdminHeader.Title>
        <AdminHeader.Description
        >Overblik over alle venues, som er associerede med events for Konnekt.</AdminHeader.Description
        >
        <AdminHeader.Actions>
          <LinkButton
            to="/admin/venues/create"
            disabled={!hasPermissions(['edit:venue'])}
          >
            <FaPlus />Tilføj
          </LinkButton>
        </AdminHeader.Actions>
      </AdminHeader>

      <main className="pt-16">
        <VenueList venues={venues} />
      </main>
    </>
  )
}
