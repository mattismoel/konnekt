import { createFileRoute, Link } from '@tanstack/react-router'
import { socialUrlToIcon, type Artist } from '@/lib/features/artist/artist';
import { pickRandom } from '@/lib/array';
import { createContext, useContext, useEffect, useRef, useState } from 'react';
import { cn } from '@/lib/clsx';
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

type ArtistsContext = {
  artists: Artist[]
  selected: Artist | undefined

  onSelect: (artist: Artist) => void;
  onExit: () => void;
}

const ArtistsContext = createContext<ArtistsContext | undefined>(undefined)

const useArtistsContext = () => {
  const ctx = useContext(ArtistsContext)
  if (!ctx) throw new Error("No ArtistsContext.Provider found")
  return ctx
}

function RouteComponent() {
  const { data: { records: artists } } = useSuspenseQuery(artistsQueryOpts)

  const [selected, setSelected] = useState<Artist>();
  const intervalRef = useRef<NodeJS.Timeout | null>(null)

  useEffect(() => {
    if (artists.length > 0) setSelected(artists[0])
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

      if (newArtist) setSelected(newArtist);
    }, 1000 / AUTO_DISPLAY_RATE);
  };

  const endAutoDisplay = () => {
    if (!intervalRef.current) return;

    clearInterval(intervalRef.current);
    intervalRef.current = null
  };

  const onSelect = (artist: Artist) => {
    setSelected(artist)
    endAutoDisplay()
  }

  return (
    <ArtistsContext.Provider value={{ artists, selected, onSelect, onExit: beginAutoDisplay }}>
      <PageMeta
        title="Konnekt | Kunstnere"
        description="Se alle aktuelle kunstnere der medvirker i Konnekts kommende events"
      />

      <main className="px-auto h-svh pt-24 md:pt-32">
        {artists.map(artist => (
          <img
            key={artist.id}
            src={artist.imageUrl}
            alt={artist.name}
            className={cn("pointer-events-none absolute top-0 left-0 -z-10 h-full w-full object-cover opacity-0 brightness-50 transition-all duration-1000", {
              "opacity-100 scale-105": selected?.id === artist.id
            })}
          />
        ))}
        <div className="space-y-16">
          <section className="flex flex-col">
            <h1 className="font-heading mb-4 text-5xl font-bold md:text-7xl text-shadow-md/15">Kunstnere</h1>
            <span className="text-text/75 text-shadow-sm">
              Her kan du se alle kunstnere, som medvirker i kommende events.
            </span>
          </section>
          {/*  ARTISTS */}
          {artists.length <= 0 && (
            <span>Der er ingen aktuelle kunstnere i øjeblikket...</span>
          )}
          <ArtistList />
        </div>
      </main>
    </ArtistsContext.Provider>
  )
}

const ArtistList = () => {
  const { artists } = useArtistsContext()

  return (
    <div className="relative">
      <ul className="max-h-96 overflow-y-scroll">
        {artists.map(artist => (
          <Entry key={artist.id} artist={artist} />
        ))}
      </ul>
    </div>
  )
}

type EntryProps = {
  artist: Artist;
}

const Entry = ({ artist }: EntryProps) => {
  const { selected, onSelect, onExit } = useArtistsContext()

  const genreString = artist.genres.map(({ name }) => name).join(", ")

  const ref = useRef<HTMLLIElement>(null)

  useEffect(() => {
    if (selected?.id === artist.id)
      ref.current?.scrollIntoView({ behavior: "smooth", block: "nearest" })
  }, [selected])

  return (
    <li
      ref={ref}
      className={cn("@container px-4 border border-transparent rounded-md hover:bg-text/10", {
        "border-text/25": selected?.id === artist.id
      })}
      onMouseEnter={() => onSelect(artist)}
      onMouseLeave={onExit}
    >
      <div className="grid grid-cols-1 @md:grid-cols-2 @2xl:grid-cols-3 items-center text-shadow-sm">
        <Link
          to="/artists/$artistId"
          params={{ artistId: artist.id.toString() }}
          className={cn("font-bold w-full py-3 text-text/50 ", {
            "text-text": selected?.id === artist.id
          })}
        >
          {artist.name}
        </Link>
        <span className="hidden @md:block text-text/75">{genreString}</span>
        <div className="hidden @2xl:flex justify-end">
          <SocialList socials={artist.socials} />
        </div>
      </div>
    </li>
  )
}

const SocialList = ({ socials }: { socials: string[] }) => {
  return (
    <ul className="flex gap-4">
      {socials.map(social => {
        const Icon = socialUrlToIcon(social)
        return (
          <li key={social} className="text-text/50">
            <a href={social}><Icon key={social} className="text-2xl" /></a>
          </li>
        )
      })}
    </ul>
  )
}
