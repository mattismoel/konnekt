import { createFileRoute, Link } from '@tanstack/react-router'
import { socialUrlToIcon, type Artist } from '@/lib/features/artist/artist';
import { pickRandom } from '@/lib/array';
import { createContext, useContext, useEffect, useRef, useState } from 'react';
import { cn } from '@/lib/clsx';
import { useSuspenseQuery } from '@tanstack/react-query';
import { previousArtistsQueryOpts, upcomingArtistsQueryOpts } from '@/lib/features/artist/query';
import PageMeta from '@/lib/components/page-meta';

/** @description The rate of which artist auto display changes artist. */
const AUTO_DISPLAY_RATE = 0.25;

/** @description The max amount of previous artists to show. */
const MAX_PREVIOUS_ARTIST_COUNT = 8

export const Route = createFileRoute('/_app/artists/')({
	component: RouteComponent,
	loader: async ({ context: { queryClient } }) => {
		queryClient.ensureQueryData(upcomingArtistsQueryOpts)
		queryClient.ensureQueryData(previousArtistsQueryOpts)
	}
})

type ArtistsContext = {
	upcomingArtists: Artist[]
	previousArtists: Artist[]
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
	const { data: upcomingArtists } = useSuspenseQuery(upcomingArtistsQueryOpts)
	const { data: previousArtists } = useSuspenseQuery(previousArtistsQueryOpts)

	const [selected, setSelected] = useState<Artist>();
	const intervalRef = useRef<NodeJS.Timeout | null>(null)

	useEffect(() => {
		if (upcomingArtists.length <= 0 && previousArtists.length > 0) setSelected(previousArtists[0])

		if (upcomingArtists.length > 0) setSelected(upcomingArtists[0])
	}, [upcomingArtists])

	useEffect(() => {
		if (upcomingArtists.length > 0 || previousArtists.length > 0) beginAutoDisplay();
		return endAutoDisplay;
	}, [upcomingArtists]);

	const beginAutoDisplay = () => {
		if (intervalRef.current) return

		intervalRef.current = setInterval(() => {
			if (upcomingArtists.length <= 0 && previousArtists.length <= 0) return
			const newArtist = pickRandom(upcomingArtists.length > 0 ? upcomingArtists : previousArtists.slice(0, MAX_PREVIOUS_ARTIST_COUNT))
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
		<ArtistsContext.Provider value={{ upcomingArtists, previousArtists, selected, onSelect, onExit: beginAutoDisplay }}>
			<PageMeta
				title="Konnekt | Kunstnere"
				description="Se alle aktuelle kunstnere der medvirker i Konnekts kommende events"
			/>

			<main className="px-auto min-h-svh pb-32 pt-24 md:pt-32">
				{[...upcomingArtists, ...previousArtists].map(artist => (
					<img
						key={artist.id}
						src={artist.imageUrl}
						alt={artist.name}
						className={cn("pointer-events-none fixed top-0 left-0 -z-10 h-lvh w-full object-cover opacity-0 brightness-50 transition-all duration-1000", {
							"opacity-100 scale-105": selected?.id === artist.id
						})}
					/>
				))}

				<div className="h-full flex flex-col">
					<section className="flex flex-col mb-8">
						<h1 className="font-heading mb-4 text-5xl font-bold md:text-7xl text-shadow-md/15">Kunstnere</h1>
						{upcomingArtists.length > 0 && (
							<span className="text-text/75 text-shadow-sm leading-relaxed">
								Her kan du se alle kunstnere, som medvirker i kommende events samt dem, der har været en del af tidligere events.
							</span>
						)}
					</section>

					<ArtistList />
				</div>
			</main>
		</ArtistsContext.Provider>
	)
}

const ArtistList = () => {
	const { upcomingArtists, previousArtists } = useArtistsContext()

	return (
		<div className="flex flex-col gap-16">
			<section>
				{/* <h2 className='mb-4 font-semibold font-heading'>Kommende</h2> */}
				{upcomingArtists.length > 0 ? (
					<ul className="flex-1 overflow-y-scroll">
						{upcomingArtists.map(artist => (
							<Entry key={artist.id} artist={artist} />
						))}
					</ul>
				) : (
					<p className="text-text/50 italic">Der er ingen kommende kunstnere...</p>
				)}
			</section>

			{previousArtists.length > 0 && (
				<section>
					<h2 className="mb-4 font-semibold font-heading">Tidligere kunstnere</h2>
					<ul className="flex-1 overflow-y-scroll">
						{previousArtists.slice(0, MAX_PREVIOUS_ARTIST_COUNT).map(artist => (
							<Entry key={artist.id} artist={artist} />
						))}
					</ul>
				</section>
			)}
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
