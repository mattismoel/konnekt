import { useSuspenseQuery } from "@tanstack/react-query"
import { createFileRoute } from "@tanstack/react-router"
import { spotifyTrackIdFromUrl } from "@/lib/spotify"
import { artistByIdQueryOpts } from "@/lib/features/artists/query"
import { artistEventsQueryOpts } from "@/lib/features/events/query"
import type { Artist } from "@/lib/features/artists/artist"
import SlideshowGallery from "@/lib/components/slideshow-gallery"
import SocialList from "@/lib/components/social-list"
import SpotifyPreview from "@/lib/components/spotify-preview"
import EventCard from "@/lib/features/events/components/event-card"

export const Route = createFileRoute("/_app/artists/$artistId")({
	component: RouteComponent,
	loader: async ({ params: { artistId }, context: { queryClient } }) => {
		const id = parseInt(artistId)
		queryClient.ensureQueryData(artistByIdQueryOpts(id))
		queryClient.ensureQueryData(artistEventsQueryOpts(id))
		return { id }
	}
})

function RouteComponent() {
	const { id } = Route.useLoaderData()
	const { data: artist } = useSuspenseQuery(artistByIdQueryOpts(id))
	const { data: { records: events } } = useSuspenseQuery(artistEventsQueryOpts(id))

	const trackId = spotifyTrackIdFromUrl(artist.previewUrl)

	return (
		<main className="min-h-svh">
			<Header artist={artist} />
			<article className="px-responsive py-16 flex flex-col gap-8">
				<section className="paragraphs-relaxed" dangerouslySetInnerHTML={{ __html: artist.description }} />
				<section>
					{trackId && (<SpotifyPreview trackId={trackId} />)}

					{events.length > 0 && (
						<>
							<h2 className="text-4xl font-semibold font-heading text-heading mb-8">Oplev {artist.name} her</h2>
							<SlideshowGallery>
								{events.map(event => <EventCard event={event} />)}
							</SlideshowGallery>
						</>
					)}
				</section>
			</article>
		</main>
	)
}

type HeaderProps = {
	artist: Artist
}

const Header = ({ artist }: HeaderProps) => {
	return (
		<header className="relative h-[80svh]">
			<img src={artist.imageUrl} className="h-full w-full object-cover" />
			<div className="fade-to-b-96 fade-background" />
			<div className="absolute bottom-16 left-0 px-responsive">
				<h1 className="text-7xl font-bold font-heading text-heading mb-8 word-spacing-[100vw]">{artist.name}</h1>
				<SocialList size="lg" urls={artist.socials} />
			</div>
		</header>
	)
}
