import { createFileRoute, Link } from '@tanstack/react-router'
import { socialUrlToIcon, type Artist } from '@/lib/features/artist';
import { pickRandom } from '@/lib/array';
import { useEffect, useRef, useState } from 'react';
import { cn } from '@/lib/clsx';
import { useListUpcomingArtists } from '@/lib/features/hook';

/** @description The rate of which artist auto display changes artist. */
const AUTO_DISPLAY_RATE = 0.25;

export const Route = createFileRoute('/_app/artists/')({
  component: RouteComponent,
})

function RouteComponent() {
  const { data: artists, isLoading } = useListUpcomingArtists()

  const [selected, setSelected] = useState<Artist>();
  const intervalRef = useRef<NodeJS.Timeout | null>(null)

  useEffect(() => {
    if (!artists) return
    const initialArtist = artists.at(0)

    if (!initialArtist) return
    setSelected(artists?.at(0))
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

  if (isLoading) return <p>Loading...</p>

  if (!artists || artists.length <= 0) return <span>Ingen kunstnere...</span>

  return (
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
        <ul
          className="divide-text/50 max-h-96 divide-y overflow-y-scroll"
          onMouseLeave={() => beginAutoDisplay()}
          onMouseEnter={() => endAutoDisplay()}
        >
          {artists.map(artist => (
            <Entry selected={selected} key={artist.id} artist={artist} onSelect={() => setSelected(artist)} />
          ))}
        </ul>
      </div>
    </main>
  )
}

type EntryProps = {
  artist: Artist;
  selected?: Artist;
  onSelect: () => void;
}

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
