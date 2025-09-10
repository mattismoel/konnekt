import { createFileRoute } from '@tanstack/react-router'

import OdenseKommuneLogo from '@/lib/assets/logos/odense-kommune-logo.svg';
import UngOdenseLogo from '@/lib/assets/logos/ungodense-logo.svg';
import PostenLogo from '@/lib/assets/logos/posten-logo.svg';
import KulturMaskinenLogo from '@/lib/assets/logos/kulturmaskinen-logo.svg';
import SpillestedetOdenseLogo from '@/lib/assets/logos/spillestedet-odense-logo.svg';
import StormsPakhusLogo from "@/lib/assets/logos/storms-pakhus-logo.svg"
import NoctivagaLogo from "@/lib/assets/logos/noctivaga-logo.svg"

import Fader from '@/lib/components/fader';
import GlowCursor from '@/lib/components/glow-cursor';
import { FaArrowRight } from 'react-icons/fa';

import LinkButton from '@/lib/components/ui/button/link-button';
import { useSuspenseQuery } from '@tanstack/react-query';
import { upcomingEventsQueryOpts } from '@/lib/features/event/query';
import PageMeta from '@/lib/components/page-meta';
import EventGrid from '@/lib/features/event/components/event-grid';
import Slideshow from '@/lib/components/slideshow';
import { landingImagesQueryOptions } from '@/lib/features/content/query';
import TeamDisplay from '@/lib/components/team-display';
import { membersQueryOpts, teamsQueryOpts } from '@/lib/features/auth/query';
import SponsorDisplay from '@/lib/components/sponsor-display';

export const Route = createFileRoute('/_app/')({
	component: App,
	loader: async ({ context: { queryClient } }) => {
		queryClient.ensureQueryData(upcomingEventsQueryOpts())
		queryClient.ensureQueryData(landingImagesQueryOptions)
		queryClient.ensureQueryData(teamsQueryOpts)
		queryClient.ensureQueryData(membersQueryOpts)
	}
})

function App() {
	const { data: { records: upcomingEvents } } = useSuspenseQuery(upcomingEventsQueryOpts())
	const { data: landingImages } = useSuspenseQuery(landingImagesQueryOptions)
	const { data: { records: teams } } = useSuspenseQuery(teamsQueryOpts)
	const { data: { records: members } } = useSuspenseQuery(membersQueryOpts)

	console.log(members)

	return (
		<>
			<PageMeta
				title="Konnekt | Forside"
				description="Forsiden til Konnekts.
        Her kan du finde kommende events, aktuelle kunstnere mm."
			/>

			<div>
				<section className="px-auto -z-50 flex h-svh flex-col justify-center gap-16">
					<div
						className="pointer-events-none absolute top-0 left-0 isolate -z-10 h-full w-full overflow-hidden"
					>
						<GlowCursor />
						<Slideshow
							srcs={landingImages.map(({ url }, index) => ({
								src: url,
								alt: `Baggrundsbillede ${index + 1}`
							}))}
						/>
					</div>
					<section className="flex max-w-lg flex-col gap-16 overflow-hidden">
						<div className="flex flex-col gap-4">
							<h2 className="font-heading text-4xl sm:text-5xl text-shadow-lg/15"><b>Fynsk musik</b> med fremtiden for øje</h2>
							<p className="text-text/75 text-shadow-md leading-relaxed">
								Et springbræt for unge, aspirerende musiskere, og en indgang ind til den danske musikscene.
							</p>
						</div>

						<div className="z-10 flex w-full flex-col-reverse gap-4 sm:flex-row">
							<LinkButton to="/about" variant="secondary" className="w-full sm:w-fit">Læs mere</LinkButton>
							<LinkButton to={upcomingEvents.length > 0 ? "/events" : "/artists"} className="group w-full items-center gap-2 sm:w-fit">
								{upcomingEvents.length > 0 ? "Se events" : "Se kunstnere"}
								<FaArrowRight className="text-sm transition-transform group-hover:translate-x-1" />
							</LinkButton>
						</div>
					</section>
					<Fader direction="up" className="absolute h-64 from-black/75 z-0" />
				</section>

				<section className="bg-zinc-950">

					<section className="px-auto flex flex-col gap-32 py-16">
						<section>
							<h1 className="font-heading mb-8 text-2xl font-bold">
								Vores mission
							</h1>
							<p className="text-text/75 leading-loose">
								Konnekt er en ungedrevet forening og et koncertinitiativ, der
								arbejder for at give unge musikere mulighed for at komme på
								scenen og få kontakt til et publikum. Vi arrangerer koncerter,
								hvor nye artister spiller sammen med mere etablerede
								upcoming-navne, og hvor publikum får adgang til musik, de ellers
								ikke ville have mødt.
								<br />
								<br />
								Men Konnekt handler ikke kun om musik – det handler om
								fællesskab og deltagelse. Vi skaber rum for unge, der vil være
								med som frivillige, arrangører eller idéfolk, og vi samarbejder
								med lokale aktører for at skabe synlighed, netværk og udvikling
								i Odenses kulturliv.
								<br />
								<br />
								Hos os er alle med til at forme oplevelsen – både på og bag
								scenen.
							</p>
						</section>

						<section className="z-0 flex w-full flex-col gap-8">
							<span className="text-center text-text/50">I samarbejde med</span>
							<div className="relative isolate w-full">
								<SponsorDisplay
									srcs={new Map([
										['https://postenlive.dk/', { src: PostenLogo, alt: "Posten" }],
										["https://stormspakhus.dk/", { src: StormsPakhusLogo, alt: "Storms Pakhus" }],
										['https://ungodense.dk/', { src: UngOdenseLogo, alt: "UngOdense" }],
										['https://odense.dk', { src: OdenseKommuneLogo, alt: "Odense Kommune" }],
										['https://ungodense.dk/index.php?open=1283&menu_id=58', { src: SpillestedetOdenseLogo, alt: "Spillestedet" }],
										['https://kulturmaskinen.dk/', { src: KulturMaskinenLogo, alt: "Kulturmaskinen" }],
										['https://www.noctivaga.dk/', { src: NoctivagaLogo, alt: "Noctivaga" }],
									])}
								/>
							</div>
						</section>
						<section>
							<h1 className="font-heading text-4xl font-bold text-center mb-16">Mød holdet</h1>
							<TeamDisplay allTeams={teams} members={members} />
						</section>

						{upcomingEvents.length > 0 && (
							<section>
								<h1 className="font-heading mb-8 text-4xl font-bold">Ses vi her?</h1>
								<EventGrid events={[upcomingEvents[0]]} />
							</section>
						)}
					</section>
				</section>
			</div >
		</>
	)
}
