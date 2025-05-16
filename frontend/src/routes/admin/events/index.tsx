import { useAuth } from '@/lib/context/auth'
import { createFileRoute } from '@tanstack/react-router'

import { useListPreviousEvents, useListUpcomingEvents } from '@/lib/features/hook';
import AdminHeader from '@/lib/components/admin-header';
import LinkButton from '@/lib/components/ui/button/link-button';
import { FaPlus } from 'react-icons/fa6';
import EventList from '@/lib/components/event-list';

export const Route = createFileRoute('/admin/events/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { hasPermissions } = useAuth()

  const upcomingEventsQuery = useListUpcomingEvents()
  const previousEventsQuery = useListPreviousEvents()

  const isLoading = upcomingEventsQuery.isLoading || previousEventsQuery.isLoading
  const isError = upcomingEventsQuery.isError || previousEventsQuery.isError

  if (isLoading) return <p>Loading...</p>
  if (isError) return <p>Error...</p>

  const upcomingEvents = upcomingEventsQuery.data?.records || []
  const previousEvents = previousEventsQuery.data?.records || []

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
