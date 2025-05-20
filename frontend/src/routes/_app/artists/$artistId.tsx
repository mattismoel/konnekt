import Caroussel from '@/lib/components/caroussel';
import Fader from '@/lib/components/fader';
import PageMeta from '@/lib/components/page-meta';
import SpotifyPreview from '@/lib/components/spotify-preview';
import { socialUrlToIcon } from '@/lib/features/artist/artist';
import { createArtistByIdOpts, createArtistEventsOpts } from '@/lib/features/artist/query';
import EventCard from '@/lib/features/event/components/event-card';
import { trackIdFromUrl } from '@/lib/spotify';
import { useSuspenseQuery } from '@tanstack/react-query';
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/artists/$artistId')({
component: RouteComponent,
loader: async ({context: {queryClient}, params: {artistId: artistIdStr}}) => {
  const artistId = parseInt(artistIdStr)
  const artistOptions = createArtistByIdOpts(artistId)
  const artistEventsOpts = createArtistEventsOpts(artistId)

queryClient.ensureQueryData(artistEventsOpts)
queryClient.ensureQueryData(artistOptions)

return {artistOptions, artistEventsOpts}
}
})

function RouteComponent() {
const {artistOptions, artistEventsOpts} = Route.useLoaderData()

const {data: artist} = useSuspenseQuery(artistOptions)
const {data: {records: artistEvents}} = useSuspenseQuery(artistEventsOpts)

if (!artist) return <p>No such artist...</p>

let trackId = artist.previewUrl ? trackIdFromUrl(artist.previewUrl) : undefined;

return  (
    <>
    <PageMeta 
      title={`Konnekt | Kunstner | ${artist.name}`}
      description={`Oplev ${artist.name} til "${artistEvents.at(0)?.title}"`}
    />

<main>
<div className="grid min-h-svh grid-cols-1 grid-rows-[85svh_1fr]">
<div className="px-auto relative isolate flex items-end py-16">
<img
src={artist.imageUrl}
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
<section className="space-y-8">
<div className="prose prose-lg md:prose-base prose-invert max-w-none" dangerouslySetInnerHTML={{ __html: artist.description}}/>
{trackId && (
<SpotifyPreview trackId={trackId} />
)}
</section>

{artistEvents.length > 0 && (
<section>
<h1 className="font-heading mb-8 text-2xl font-bold">Oplev {artist.name} her</h1>
<Caroussel>
{artistEvents.map(event => <EventCard key={event.id} event={event}/>)}
</Caroussel>
</section>
)}
</article>
</div>
</main>
    </>
)
}

