import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/events/edit')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/admin/events/edit"!</div>
}
