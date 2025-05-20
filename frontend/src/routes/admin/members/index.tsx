import AdminHeader from '@/lib/components/admin-header'
import MemberList from '@/lib/features/auth/components/member-list'
import { membersQueryOpts, pendingMembersQueryOpts } from '@/lib/features/auth/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/members/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(membersQueryOpts)
    queryClient.ensureQueryData(pendingMembersQueryOpts)
  }
})

function RouteComponent() {
  const { data: { records: members } } = useSuspenseQuery(membersQueryOpts)
  const { data: { records: pendingMembers } } = useSuspenseQuery(pendingMembersQueryOpts)

  return (
    <>
      <AdminHeader>
        <AdminHeader.Title>Medlemmer</AdminHeader.Title>
        <AdminHeader.Description
        >Her kan du administrere alle medlemmer af foreningen, herunder deres roller i foreningen, og
          dermed deres tilladelser.</AdminHeader.Description
        >
      </AdminHeader>

      <main className="pt-16">
        <MemberList members={members} pendingMembers={pendingMembers} />
      </main>
    </>
  )
}
