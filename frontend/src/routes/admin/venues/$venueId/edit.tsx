import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { createVenueByIdQueryOptions } from '@/lib/features/event/query'

import VenueForm from '@/lib/features/event/components/venue-form'
import { createMemberByIdQueryOpts } from '@/lib/features/auth/query'

export const Route = createFileRoute('/admin/venues/$venueId/edit')({
	component: RouteComponent,
	loader: async ({ context: { queryClient }, params: { venueId } }) => {
		const venueQueryOptions = createVenueByIdQueryOptions(parseInt(venueId))
		const venue = await queryClient.ensureQueryData(venueQueryOptions)

		const updatedByMemberQueryOpts = createMemberByIdQueryOpts(venue.updatedBy)

		return { venueQueryOptions, updatedByMemberQueryOpts }
	}
})

function RouteComponent() {
	const { venueQueryOptions, updatedByMemberQueryOpts } = Route.useLoaderData()
	const { data: venue } = useSuspenseQuery(venueQueryOptions)
	const { data: updatedByMember } = useSuspenseQuery(updatedByMemberQueryOpts)

	return (
		<main className="px-auto py-32 min-h-svh flex justify-center items-center">
			<div>
				<h1 className="font-heading text-2xl font-bold mb-4">Redigér venue</h1>
				<VenueForm venue={venue} updatedByMember={updatedByMember} />
			</div>
		</main>
	)
}
