import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/events/')({
	component: RouteComponent,
})

function RouteComponent() {
	return (
		<main className='min-h-svh py-32 px-responsive'>
			Events page.
		</main>
	)
}
