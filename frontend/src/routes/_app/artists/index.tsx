import { createFileRoute, Link } from '@tanstack/react-router'
import { socialUrlToIcon, type Artist } from '@/lib/features/artist/artist';
import { pickRandom } from '@/lib/array';
import { useEffect, useRef, useState } from 'react';
import { cn } from '@/lib/clsx';
import { randomInt } from '@/lib/random';
import { SkeletonIcon, SkeletonText } from '@/lib/components/skeleton';
import { useSuspenseQuery } from '@tanstack/react-query';
import { artistsQueryOpts } from '@/lib/features/artist/query';
import PageMeta from '@/lib/components/page-meta';

/** @description The rate of which artist auto display changes artist. */
const AUTO_DISPLAY_RATE = 0.25;

export const Route = createFileRoute('/_app/artists/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(artistsQueryOpts)
  }
})

function RouteComponent() {
  const { data: { records: artists } } = useSuspenseQuery(artistsQueryOpts)

  const [selected, setSelected] = useState<Artist>();
  const intervalRef = useRef<NodeJS.Timeout | null>(null)

  useEffect(() => {
    if (artists.length <= 0) return

    const initialArtist = artists[0]
    setSelected(initialArtist)
  }, [artists])

  useEffect(() => {
    beginAutoDisplay();
    return endAutoDisplay;
  }, [artists]);

  const beginAutoDisplay = () => {
    if (intervalRef.current) return

    intervalRef.current = setInterval(() => {
      if (!artists || artists.length <= 0) return

      const newArtist = pickRandom(artists, selected);

      if (!newArtist) return

      setSelected(newArtist);
    }, 1000 / AUTO_DISPLAY_RATE);
  };

  const endAutoDisplay = () => {
    if (!intervalRef.current) return;

    clearInterval(intervalRef.current);
    intervalRef.current = null
  };

  return (
    <>
      <PageMeta
        title="Konnekt | Kunstnere"
        description="Se alle aktuelle kunstnere der medvirker i Konnekts kommende events"
      />

      <main className="px-auto h-svh pt-32">
        {artists.map(artist => (
          <img
            key={artist.id}
            src={artist.imageUrl}
            alt={artist.name}
            className={cn("pointer-events-none absolute top-0 left-0 -z-10 h-full w-full object-cover opacity-0 brightness-75 transition-all duration-1000", {
              "opacity-100 scale-105": selected?.id === artist.id
            })}
          />
        ))}
        <div className="space-y-16">
          <section className="flex flex-col">
            <h1 className="font-heading mb-4 text-5xl font-bold md:text-7xl">Kunstnere</h1>
            <span className="text-text/75">
              Her kan du se alle kunstnere, som medvirker i kommende events.
            </span>
          </section>
          {/*  ARTISTS */}
          {artists.length <= 0 && (
            <span>Der er ingen aktuelle kunstnere i øjeblikket...</span>
          )}
          <ArtistList
            artists={artists}
            selected={selected}
            onSelect={setSelected}
            onMouseEnter={endAutoDisplay}
            onMouseLeave={beginAutoDisplay}
          />
        </div>
      </main>
    </>
  )
}

type EntryProps = {
  artist: Artist;
  selected?: Artist;
  onSelect: () => void;
}

type ArtistListProps = {
  artists: Artist[] | undefined
  selected: Artist | undefined

  onMouseLeave: () => void;
  onMouseEnter: () => void;

  onSelect: (artist: Artist) => void;
}

const ArtistList = ({ artists, selected, onSelect, onMouseEnter, onMouseLeave }: ArtistListProps) => (
  <ul
    className="divide-text/50 max-h-96 divide-y overflow-y-scroll"
    onMouseLeave={onMouseLeave}
    onMouseEnter={onMouseEnter}
  >
    {!artists
      ? [...Array(randomInt(4, 8))].map((_, i) => <SkeletonEntry key={i} />)
      : artists.map(artist => (
        <Entry selected={selected} key={artist.id} artist={artist} onSelect={() => onSelect(artist)} />
      )
      )}
  </ul>
)

const Entry = ({ artist, selected, onSelect }: EntryProps) => {
  return (
    <li
      className={cn("group text-text/75 hover:text-text [.selected]:text-text relative flex items-center border-l-transparent transition-colors", {
        "text-text": selected?.id === artist.id
      })}
      onMouseEnter={onSelect}
    >
      {/* <!-- SELECTED MARKER --> */}
      <div
        className="group-[.selected]:bg-text h-6 w-1 scale-y-0 rounded-full bg-transparent transition-all group-[.selected]:scale-y-100"
      ></div>
      <div className="grid w-full grid-cols-3">
        <Link
          to="/artists/$artistId"
          params={{
            artistId: artist.id.toString(),
          }}
          className="col-span-2 grid grid-cols-2 py-3 pl-3"
        >
          <span className="line-clamp-1 font-bold">{artist.name}</span>
          <span className="line-clamp-1">{artist.genres.map((g) => g.name).join(', ')}</span>
        </Link>
        <div
          className="text-text/50 group-[.selected]:text-text/75 group-hover:text-text/75 flex items-center justify-end gap-2 pr-3 text-lg"
        >
          {artist.socials.map(social => {
            const Icon = socialUrlToIcon(social)

            return (
              <a key={social} href={social} className="hover:text-text">
                <Icon />
              </a>
            )
          })}
        </div>
      </div>
    </li>
  )
}

const SkeletonEntry = () => (
  <li
    className="group text-text/75 relative flex items-center"
  >
    <div className="grid w-full grid-cols-3 py-3 pl-3 items-center">
      {/* ARTIST NAME */}
      <SkeletonText />

      {/* GENRES */}
      <SkeletonText wordCount={randomInt(2, 4)} />

      {/* SOCIALS */}
      <div className="text-text/50 flex items-center justify-end gap-2 pr-3">
        <SkeletonIcon />
      </div>
    </div>
  </li>
)

