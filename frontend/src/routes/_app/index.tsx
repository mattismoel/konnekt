import { createFileRoute } from '@tanstack/react-router'
import { useEffect, type HTMLAttributes } from 'react';
import { useSuspenseQuery } from '@tanstack/react-query';
import { upcomingEventsQueryOpts } from '@/lib/features/events/query';
import { membersQueryOpts } from '@/lib/features/auth/query';
import { useMousePos } from '@/lib/hooks/useMousePos';
import { useRandomIndex } from '@/lib/hooks/useRandom';
import { cn } from '@/lib/clsx';

import Posten from "@/lib/assets/sponsors/posten-logo.svg"
import StormsPakhus from "@/lib/assets/sponsors/storms-pakhus-logo.svg"
import UngOdense from "@/lib/assets/sponsors/ungodense-logo.svg"
import OdenseKommune from "@/lib/assets/sponsors/odense-kommune-logo.svg"
import SpillestedetOdense from "@/lib/assets/sponsors/spillestedet-odense-logo.svg"
import Kulturmaskinen from "@/lib/assets/sponsors/kulturmaskinen-logo.svg"
import Noctivaga from "@/lib/assets/sponsors/noctivaga-logo.svg"

import GlowingCursor from '@/lib/components/glowing-cursor';
import SlideshowGallery from '@/lib/components/slideshow-gallery';
import LinkButton from '@/lib/components/ui/button/link-button';
import LogoDisplay from '@/lib/components/logo-display';
import TeamDisplay from '@/lib/components/member-display';
import EventCard from '@/lib/features/events/components/event-card';

export const Route = createFileRoute('/_app/')({
	component: App,
	loader: async ({ context: { queryClient } }) => {
		queryClient.ensureQueryData(upcomingEventsQueryOpts())
		queryClient.ensureQueryData(membersQueryOpts())
	}
})

const images: string[] = [...Array(5)].map((_, i) => `https://picsum.photos/seed/image-${i}/4096`)

function App() {
	const { data: { records: upcomingEvents } } = useSuspenseQuery(upcomingEventsQueryOpts())
	const { data: { records: members } } = useSuspenseQuery(membersQueryOpts())

	const { randomIndex, randomize } = useRandomIndex(images)
	const mousePos = useMousePos()

	useEffect(() => {
		const interval = setInterval(() => randomize(), 5000)
		return () => clearInterval(interval)
	}, [])

	return (
		<main>
			<section className='relative isolate px-responsive flex flex-col justify-center gap-4 h-svh overflow-hidden'>
				{images.map((src, idx) => (
					<img key={src} src={src} className={cn('h-full w-full z-0 absolute top-0 left-0 brightness-40 object-cover opacity-0 transition-opacity duration-2000 [.active]:opacity-100', idx === randomIndex && "active")} />
				))}

				<GlowingCursor size="xl" mousePos={mousePos} />

				<div className='z-10 flex flex-col gap-16'>
					<div>
						<h1 className='text-5xl text-heading font-heading mb-4'>
							<b>Fynsk musik</b> med<br />
							fremtiden for øje
						</h1>
						<p className='max-w-md text-heading/75'>
							Et springbræt for unge, aspirerende musiskere, og en indgang ind til
							den danske musikscene.
						</p>
					</div>

					<div className='flex flex-col gap-4 sm:flex-row sm:w-fit'>
						<LinkButton to="/about-us" variant='secondary'>Læs mere</LinkButton>
						<LinkButton to="/events" variant='primary'>Find billetter</LinkButton>
					</div>
				</div>
			</section>

			{/* MISSION STATEMENT  */}
			<InfoSection className="paragraphs-relaxed">
				<h2 className='text-2xl font-semibold mb-4 text-heading font-heading'>Vores mission</h2>
				<p className='leading-relaxed'>
					Konnekt er en ungedrevet forening og et koncertinitiativ, der arbejder
					for at give unge musikere mulighed for at komme på scenen og få
					kontakt til et publikum. Vi arrangerer koncerter, hvor nye artister
					spiller sammen med mere etablerede upcoming-navne, og hvor publikum
					får adgang til musik, de ellers ikke ville have mødt.
				</p>
				<p>
					Men Konnekt handler ikke kun om musik – det handler om fællesskab og
					deltagelse. Vi skaber rum for unge, der vil være med som frivillige,
					arrangører eller idéfolk, og vi samarbejder med lokale aktører for at
					skabe synlighed, netværk og udvikling i Odenses kulturliv.
				</p>
				<p>
					Hos os er alle med til at forme oplevelsen – både på og bag scenen.
				</p>
			</InfoSection>

			{/* SPONSOR DISPLAY */}
			<InfoSection>
				<h3 className='text-center mb-4 text-text/75'>I samarbejde med</h3>
				<LogoDisplay
					srcs={[
						{ name: "Posten", href: "https://postenlive.dk", src: Posten },
						{ name: "Storms Pakhus", href: "https://stormspakhus.dk", src: StormsPakhus },
						{ name: "UngOdense", href: "https://ungodense.dk", src: UngOdense },
						{ name: "Odense Kommune", href: "https://odense.dk", src: OdenseKommune },
						{ name: "Spillestedet Odense", href: "https://ungodense.dk/index.php?open=1283&menu_id=58", src: SpillestedetOdense },
						{ name: "Kulturmaskinen", href: "https://kulturmaskinen.dk", src: Kulturmaskinen },
						{ name: "Noctivata", href: "https://noctivaga.dk", src: Noctivaga },
					]} />
			</InfoSection>

			{/* TEAM DISPLAY */}
			<InfoSection>
				<h2 className='text-4xl font-heading font-semibold text-heading mb-12 text-center'>Mød holdet</h2>
				<TeamDisplay members={members} />
			</InfoSection>

			{/* UPCOMING EVENTS DISPLAY */}
			{upcomingEvents.length > 0 && (
				<InfoSection>
					<h2 className="text-4xl font-heading font-semibold text-heading mb-8">Ses vi her?</h2>
					<SlideshowGallery>
						{upcomingEvents.map(event => <EventCard key={event.id} event={event} />)}
					</SlideshowGallery>
				</InfoSection>
			)}
		</main >
	)
}

const InfoSection = ({ children, className, ...rest }: HTMLAttributes<HTMLDivElement>) => (
	<section {...rest} className={cn("bg-background px-responsive py-16 first:pt-32 last:pb-32", className)}>
		{children}
	</section>
)
