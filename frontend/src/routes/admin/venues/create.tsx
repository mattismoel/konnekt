import { createFileRoute } from '@tanstack/react-router'

import VenueForm from '@/lib/features/event/components/venue-form'

export const Route = createFileRoute('/admin/venues/create')({
	component: RouteComponent,
})

function RouteComponent() {

	return (
		<main className="px-auto py-32 min-h-svh flex justify-center items-center">
			<div>
				<h1 className="font-heading text-2xl font-bold mb-4">Lav venue</h1>
				<VenueForm />
			</div>
		</main >
	)
}
