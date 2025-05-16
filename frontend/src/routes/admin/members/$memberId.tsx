import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/members/$memberId')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/admin/members/$memberId"!</div>
}
