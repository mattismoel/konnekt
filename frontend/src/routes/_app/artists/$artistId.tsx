import Caroussel from '@/lib/components/caroussel';
import EventCard from '@/lib/components/event-card';
import Fader from '@/lib/components/fader';
import SpotifyPreview from '@/lib/components/spotify-preview';
import { socialUrlToIcon } from '@/lib/features/artist';
import { useArtistById, useArtistEvents } from '@/lib/features/hook';
import { trackIdFromUrl } from '@/lib/spotify';
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/artists/$artistId')({
  component: RouteComponent,
})

function RouteComponent() {
  const {artistId} = Route.useParams()

  const artistQuery = useArtistById(parseInt(artistId))
  const artistEventsQuery = useArtistEvents(parseInt(artistId))

  const isLoading = artistQuery.isLoading || artistEventsQuery.isLoading
  const isError = artistQuery.isError || artistEventsQuery.isError

  if (isLoading) return <p>Loading...</p>
  if (isError) return <p>Error...</p>


  const artist = artistQuery.data
  const events = artistEventsQuery.data?.records || []

  if (!artist) return <p>No such artist...</p>

	let trackId = artist.previewUrl ? trackIdFromUrl(artist.previewUrl) : undefined;

  return  (
<main>
	<div className="grid min-h-svh grid-cols-1 grid-rows-[85svh_1fr]">
		<div className="px-auto relative isolate flex items-end py-16">
			<img
				src={artistQuery.data.imageUrl}
				alt="Cover af {artist.name}"
				className="absolute top-0 left-0 h-full w-full object-cover"
			/>
			<Fader direction="right" className="absolute hidden w-96 from-zinc-950 md:block" />
			<Fader direction="up" className="absolute h-[512px] from-zinc-950" />
			<div
				className="z-10 flex w-full flex-col items-start justify-between gap-8 md:flex-row md:items-end"
			>
				<h1
					style:word-spacing="100vw"
					className="font-heading text-7xl font-bold md:text-8xl lg:text-9xl"
				>
					{artist.name}
				</h1>
				<div className="text-text/50 flex gap-4 text-3xl">
          {artist.socials.map(social => {
            const Icon = socialUrlToIcon(social)
                return (
                  <a key={social} href={social} className="hover:text-text transition-colors">
                    <Icon />
                  </a>
                )
            })}
				</div>
			</div>
		</div>
		<article className="px-auto space-y-16 bg-zinc-950 py-16">
			{/* <!-- ARTICLE CONTENT --> */}
			<section className="space-y-8">
				<div className="prose prose-lg md:prose-base prose-invert max-w-none" dangerouslySetInnerHTML={{ __html: artist.description}}/>
              
				{trackId && (
					<SpotifyPreview trackId={trackId} />
        )}
			</section>

          {events.length > 0 && (
				<section>
					<h1 className="font-heading mb-8 text-2xl font-bold">Oplev {artist.name} her</h1>
					<Caroussel>
                {events.map(event => <EventCard key={event.id} event={event}/>)}
					</Caroussel>
				</section>
          )}
		</article>
	</div>
</main>
  )
}
