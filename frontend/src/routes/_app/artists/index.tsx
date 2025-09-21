import { cn } from '@/lib/clsx'
import SocialIcon from '@/lib/components/social-icon'
import SocialList from '@/lib/components/social-list'
import type { Artist } from '@/lib/features/artists/artist'
import { previousArtistsQueryOpts, upcomingArtistsQueryOpts } from '@/lib/features/artists/query'
import { useRandomIndex } from '@/lib/hooks/useRandom'
import { socialIconByUrl } from '@/lib/social'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute, Link } from '@tanstack/react-router'
import { forwardRef, useEffect, useState, type HTMLAttributes } from 'react'
import { FaInstagram } from 'react-icons/fa6'

export const Route = createFileRoute('/_app/artists/')({
	component: RouteComponent,
	loader: async ({ context: { queryClient } }) => {
		queryClient.ensureQueryData(upcomingArtistsQueryOpts())
		queryClient.ensureQueryData(previousArtistsQueryOpts())
	}
})

function RouteComponent() {
	const { data: { records: upcomingArtists } } = useSuspenseQuery(upcomingArtistsQueryOpts())
	const { data: { records: previousArtists } } = useSuspenseQuery(previousArtistsQueryOpts())

	const { randomIndex, randomize, overrideIndex } = useRandomIndex(upcomingArtists)

	const [isManualShowcase, setIsManualShowcase] = useState(false)

	useEffect(() => {
		if (isManualShowcase) return
		const interval = setInterval(() => randomize(), 2000)
		return () => clearInterval(interval)
	}, [isManualShowcase])

	const handleSetShowcaseArtist = (idx: number) => {
		overrideIndex(idx)
		setIsManualShowcase(true)
	}

	return (
		<main className="min-h-svh py-32 px-responsive overflow-hidden">
			<div className="z-0 fixed top-0 left-0 w-full h-svh overflow-hidden">
				{[...upcomingArtists].map((artist, i) => (
					<img key={artist.imageUrl} src={artist.imageUrl} className={cn("absolute top-0 left-0 h-full w-full object-cover opacity-0 brightness-60 transition-[opacity,scale] duration-1000 [.active]:opacity-100 [.active]:scale-110", i === randomIndex && "active")} />
				))}
			</div>

			<ul className="@container isolate z-50">
				{upcomingArtists.map((artist, i) => (
					<Entry
						artist={artist}
						isShowcased={i === randomIndex}
						onMouseOver={() => handleSetShowcaseArtist(i)}
						onMouseLeave={() => setIsManualShowcase(false)}
					/>
				))}
			</ul>
		</main>
	)
}

type EntryProps = HTMLAttributes<HTMLLIElement> & {
	artist: Artist
	isShowcased: boolean
}

const Entry = forwardRef<HTMLLIElement, EntryProps>(({ artist, isShowcased, ...rest }, ref) => {
	return (
		<li ref={ref} {...rest}>
			<div className={cn("group flex flex-row items-center text-lg  border border-transparent rounded-sm *:w-full",
				"hover:bg-foreground/10 hover:border-foreground/20 [.showcase]:border-foreground/20",
				isShowcased && "showcase",
			)}>
				<Link to="/artists/$artistId" params={{ artistId: artist.id.toString() }} className="pl-4 py-2">
					<span className="font-semibold group-hover:text-foreground group-[.showcase]:text-foreground">{artist.name}</span>
				</Link>
				<span className="hidden w-full @lg:inline">{artist.genres.join(", ")}</span>
				<SocialList size="md" urls={artist.socials} className="hidden justify-end @2xl:flex pr-6" />
			</div>
		</li>
	)
})
