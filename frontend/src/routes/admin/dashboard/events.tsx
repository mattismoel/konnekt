import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/dashboard/events')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/admin/dashboard/events"!</div>
}
