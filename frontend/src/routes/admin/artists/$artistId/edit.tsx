import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { createArtistByIdOpts, genresQueryOpts } from '@/lib/features/artist/query'

import ArtistForm from '@/lib/features/artist/components/artist-form'
import { createMemberByIdQueryOpts } from '@/lib/features/auth/query'

export const Route = createFileRoute('/admin/artists/$artistId/edit')({
	component: RouteComponent,
	loader: async ({ context: { queryClient }, params: { artistId } }) => {
		const artistQueryOptions = createArtistByIdOpts(parseInt(artistId))

		const artist = await queryClient.ensureQueryData(artistQueryOptions)
		const updatedByMemberQueryOpts = createMemberByIdQueryOpts(artist.updatedBy)

		queryClient.ensureQueryData(genresQueryOpts)
		queryClient.ensureQueryData(updatedByMemberQueryOpts)

		return { artistQueryOptions, updatedByMemberQueryOpts }
	}
})

function RouteComponent() {
	const { artistQueryOptions, updatedByMemberQueryOpts } = Route.useLoaderData()
	const { data: artist } = useSuspenseQuery(artistQueryOptions)
	const { data: { records: genres } } = useSuspenseQuery(genresQueryOpts)
	const { data: updatedByMember } = useSuspenseQuery(updatedByMemberQueryOpts)

	return (
		<ArtistForm artist={artist} genres={genres} updatedByMember={updatedByMember} />
	)
}
