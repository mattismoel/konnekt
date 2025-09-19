import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/events/$eventId')({
	component: RouteComponent,
})

function RouteComponent() {
	return (
		<main className="min-h-svh py-32 px-responsive">
			Event by ID page.
		</main>
	)
}
