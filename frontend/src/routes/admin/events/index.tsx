import { useAuth } from '@/lib/context/auth'
import { createFileRoute } from '@tanstack/react-router'

import AdminHeader from '@/lib/components/admin-header';
import LinkButton from '@/lib/components/ui/button/link-button';
import { FaPlus } from 'react-icons/fa6';
import { useSuspenseQuery } from '@tanstack/react-query';
import { previousEventsQueryOpts, upcomingEventsQueryOpts } from '@/lib/features/event/query';
import EventList from '@/lib/features/event/components/event-list';

export const Route = createFileRoute('/admin/events/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(upcomingEventsQueryOpts)
    queryClient.ensureQueryData(previousEventsQueryOpts)
  }
})

function RouteComponent() {
  const { hasPermissions } = useAuth()

  const { data: { records: upcomingEvents } } = useSuspenseQuery(upcomingEventsQueryOpts)
  const { data: { records: previousEvents } } = useSuspenseQuery(previousEventsQueryOpts)

  return (
    <>
      <AdminHeader>
        <AdminHeader.Title>Events</AdminHeader.Title>
        <AdminHeader.Description>Overblik over alle events.</AdminHeader.Description>
        <AdminHeader.Actions>
          <LinkButton
            to="/admin/events/create"
            disabled={!hasPermissions(['edit:event'])}
          >
            <FaPlus />Tilføj
          </LinkButton>
        </AdminHeader.Actions>
      </AdminHeader>

      <main className="pt-16">
        <EventList previousEvents={previousEvents} upcomingEvents={upcomingEvents} />
      </main>
    </>
  )
}
