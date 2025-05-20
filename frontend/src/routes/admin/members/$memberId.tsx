import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { createMemberByIdQueryOpts, createMemberTeamsQueryOpts, teamsQueryOpts } from '@/lib/features/auth/query'

import MemberForm from '@/lib/features/auth/components/member-form'

export const Route = createFileRoute('/admin/members/$memberId')({
  component: RouteComponent,
  loader: async ({ context: { queryClient }, params: { memberId } }) => {
    const memberQueryOpts = createMemberByIdQueryOpts(parseInt(memberId))

    queryClient.ensureQueryData(memberQueryOpts)
    queryClient.ensureQueryData(teamsQueryOpts)

    const memberTeamsQueryOpts = createMemberTeamsQueryOpts(parseInt(memberId))

    queryClient.ensureQueryData(memberTeamsQueryOpts)

    return { memberQueryOpts, memberTeamsQueryOpts }
  }
})

function RouteComponent() {
  const { memberQueryOpts, memberTeamsQueryOpts } = Route.useLoaderData()

  const { data: member } = useSuspenseQuery(memberQueryOpts)
  const { data: { records: teams } } = useSuspenseQuery(teamsQueryOpts)
  const { data: memberTeams } = useSuspenseQuery(memberTeamsQueryOpts)

  return (
    <main className="flex flex-col gap-16">
      <MemberForm member={member} teams={teams} memberTeams={memberTeams} />
    </main>
  )
}
