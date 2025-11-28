import { createFileRoute, Link } from '@tanstack/react-router'
import { socialUrlToIcon, type Artist } from '@/lib/features/artist/artist';
import { useEffect, useRef, useState } from 'react';
import { cn } from '@/lib/clsx';
import { useSuspenseQuery } from '@tanstack/react-query';
import { previousArtistsQueryOpts, upcomingArtistsQueryOpts } from '@/lib/features/artist/query';
import PageMeta from '@/lib/components/page-meta';
import Button from '@/lib/components/ui/button/button';

export const Route = createFileRoute('/_app/artists/')({
	component: RouteComponent,
	loader: async ({ context: { queryClient } }) => {
		queryClient.ensureQueryData(upcomingArtistsQueryOpts)
		queryClient.ensureQueryData(previousArtistsQueryOpts)
	}
})

function RouteComponent() {
	const { data: upcomingArtists } = useSuspenseQuery(upcomingArtistsQueryOpts)
	const { data: previousArtists } = useSuspenseQuery(previousArtistsQueryOpts)

	const [selected, setSelected] = useState<Artist>();
	const [showPrevious, setShowPrevious] = useState(false)

	useEffect(() => {
		if (upcomingArtists.length <= 0 && previousArtists.length > 0) setSelected(previousArtists[0])

		if (upcomingArtists.length > 0) setSelected(upcomingArtists[0])
	}, [upcomingArtists])


	const onSelect = (artist: Artist) => {
		setSelected(artist)
	}

	return (
		<>
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
						className={cn("scale-102 pointer-events-none fixed top-0 left-0 -z-10 h-lvh w-full object-cover opacity-0 brightness-50 transition-[scale,opacity] duration-400", {
							"opacity-100 scale-100": selected?.id === artist.id
						})}
					/>
				))}

				<div className="h-full flex flex-col">
					<section className="flex flex-col mb-8">
						<h1 className="font-heading mb-4 text-4xl font-bold text-shadow-md/15">Kommende kunstnere</h1>
					</section>

					<div className="flex flex-col gap-16">
						<section>
							{upcomingArtists.length > 0 ? (
								<ArtistList artists={upcomingArtists} onSelect={onSelect} />
							) : (
								<p className="text-text/50 italic">Der er ingen kommende kunstnere...</p>
							)}
						</section>

						{(previousArtists.length > 0) && (
							showPrevious && (
								<section className="flex flex-col gap-4">
									<h2 className="font-semibold font-heading">Tidligere kunstnere</h2>
									<ArtistList artists={previousArtists} onSelect={onSelect} />
								</section>
							)
						)}

						{previousArtists.length > 0 && (
							<Button
								type="button"
								variant="secondary"
								onClick={() => setShowPrevious(prev => !prev)}
								className="w-full py-3 rounded-md border border-transparent text-text/75 hover:border-text/15 hover:text-text hover:bg-text/15 bg-text/10"
							>
								{showPrevious ? "Skjul tidligere" : "Vis tidligere"}
							</Button>
						)}
					</div>
				</div>
			</main>
		</>
	)
}

type ArtistListProps = {
	artists: Artist[]
	onSelect: (artist: Artist) => void;
}

const ArtistList = ({ artists, onSelect }: ArtistListProps) => {
	return (
		<ul className="flex-1 overflow-y-scroll">
			{artists.map(artist => (
				<Entry
					key={artist.id}
					artist={artist}
					onSelect={() => onSelect(artist)}
				/>
			))}
		</ul>
	)
}

type EntryProps = {
	artist: Artist;
	onSelect: () => void;
}

const Entry = ({ artist, onSelect }: EntryProps) => {
	const genreString = artist.genres.map(({ name }) => name).join(", ")

	const ref = useRef<HTMLLIElement>(null)

	return (
		<li
			ref={ref}
			className="group isolate relative @container px-4 border border-transparent rounded-md overflow-hidden hover:bg-text/5 transition-colors hover:border-text/25"
			onMouseEnter={onSelect}
		>
			<div className="grid grid-cols-1 @md:grid-cols-2 @2xl:grid-cols-3 items-center text-shadow-sm">
				<Link
					to="/artists/$artistId"
					params={{ artistId: artist.id.toString() }}
					className="font-bold w-full py-3 text-text/50 transition-colors group-hover:text-text"
				>
					{artist.name}
				</Link>
				<span className="hidden @md:block text-text/75 cursor-default transition-colors group-hover:text-text">{genreString}</span>
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
