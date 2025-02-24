import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/artists/$artistId')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/artists/$artistId"!</div>
}
