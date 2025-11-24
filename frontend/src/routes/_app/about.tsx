import Accordion from '@/lib/components/accordion'
import PageMeta from '@/lib/components/page-meta'
import TeamDisplay from '@/lib/components/team-display'
import { membersQueryOpts, teamsQueryOpts } from '@/lib/features/auth/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import type { PropsWithChildren } from 'react'

export const Route = createFileRoute('/_app/about')({
	component: RouteComponent,
	loader: async ({ context: { queryClient } }) => {
		queryClient.ensureQueryData(membersQueryOpts)
		queryClient.ensureQueryData(teamsQueryOpts)
	}
})

function RouteComponent() {
	const { data: { records: members } } = useSuspenseQuery(membersQueryOpts)
	const { data: { records: teams } } = useSuspenseQuery(teamsQueryOpts)

	return (
		<>
			<PageMeta
				title="Konnekt | Om os"
				description='Konnekt er et projekt med henblik på unge musikere.
            Det danske musikmiljø er for svært at bryde igennem - især for unge
            aspirerende musikere og det skal ændres.'
			/>
			<main className='min-h-svh py-32 px-auto flex flex-col gap-32'>
				<Section>
					<h1>Hvem er vi?</h1>
					<p className="leading-loose">
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
				</Section>
				<Section>
					<h1>Hvad vil vi?</h1>
					<p className="leading-loose">
						Foreningen Konnekt arbejder for at styrke unges muligheder i
						musiklivet ved at skabe en scene, hvor nye musikere kan få erfaring,
						synlighed og møde et publikum i øjenhøjde.

						Vi tror på, at stærke kulturelle fællesskaber starter med plads og
						tillid. Derfor inviterer vi unge ind som medskabere, ikke bare
						deltagere. Gennem samarbejde, frivillighed og lokale partnerskaber
						skaber vi en model for fremtidens kulturarbejde.
					</p>
				</Section>

				<section className="flex flex-col">
					<h1 className="text-center font-heading text-4xl font-bold mb-16">Mød holdet</h1>
					<TeamDisplay allTeams={teams} members={members} />
				</section>

				<section className="flex flex-col gap-16">
					<div className="mb-8">
						<h1 className="font-heading text-4xl font-bold mb-8">Ofte stillede spørgsmål</h1>
						<p className="text-text/75 leading-relaxed">
							Her kan se diverse spørgsmål, som vi ofte stilles.
							Har du yderligere spørgsmål, så&nbsp;
							<a
								className="underline"
								href="mailto:konnekt.samarbejde@gmail.com"
							>
								send os endelig en mail.
							</a>
						</p>
					</div>

					<FAQSection title="Til kunstnere">
						<Accordion title="Hvordan bliver jeg booket?">
							Konnekt er altid på udkig efter nye, spændende navne til vores koncerter. Hvis du
							eller dit band har lyst til at spille, kan I <a href="mailto:booking.konnekt@gmail.com">sende en mail</a> med
							lidt info om jer selv, og et lydklip fra en øver, koncert eller demo til.
							Vi gennemgår løbende alle henvendelser og tager jer med i overvejelserne
							til kommende koncerter.
						</Accordion>
						<Accordion title="Kan man blive booket, hvis man ikke er fra Fyn?">
							Ja, det kan du godt! Konnekt har som kerneværdi at styrke det fynske musik
							vækstlag, men tager også gerne imod bud fra resten af landet.
						</Accordion>
						<Accordion title="Kan man blive booket, hvis man ikke har udgivet musik?">
							Ja, det kan man godt! Konnekt booker udelukkende bands/musikere der optræder
							med selvskrevet musik - også selvom du/i ikke har udgivet musik.
						</Accordion>
					</FAQSection>

					<FAQSection title="Til frivillige">
						<Accordion title="Hvordan kan jeg hjælpe?">
							Konnekt er drevet af unge kræfter og gode idéer, og vi er altid på udkig efter flere.
							Uanset om du vil arrangere, planlægge, markedsføre, økonomien, skabe visuals,
							booke bands, skabe en god atmosfære, eller noget helt andet, så er der plads til
							dig.
							<br />
							<br />
							Vi tror på, at kulturlivet i Odense bliver stærkere, når flere er med til at forme det.
							Derfor lytter vi til nye perspektiver og er altid åbne for friske inputs og kreative
							hjerner. Har du lyst til at være med, så <a href="mailto:konnekt.samarbejde@gmail.com">skriv til os</a>,
							og fortæl hvad du brænder for. Vi glæder os til at høre fra dig.
						</Accordion>
						<Accordion title="Hvad laver man som frivillig?">
							Som frivillig i Konnekts bestyrelse er du med til at forme alt fra idé til virkelighed.
							Du deltager i møder, planlægger koncerter og udvikler på, hvordan vi bedst løfter ungt
							kulturliv i Odense. Det kan være booking, PR, økonomi, samarbejder eller hvordan vi
							gør publikum gladere – alt efter hvad du brænder for.
							<br />
							<br />
							Når koncertdagen kommer, hjælper du med det praktiske: kunstnerforplejning,
							opsætning, afvikling og alt det, der får aftenen til at spille. Du får erfaring,
							indflydelse og et fællesskab med andre unge, der ligesom dig gerne vil gøre en forskel.
							<br />
							<br />
							Har du lyst til at være med? Så <a href="mailto:konnekt.samarbejde@gmail.com">send os endelig en mail</a>.
							Vi vil gerne høre fra dig!
						</Accordion>
					</FAQSection>
					<FAQSection title="Til samarbejdspartnere">
						<Accordion title="Hvordan kan vi samarbejde?">
							Konnekt er bygget på et grundlag af stærke samarbejdspartnere, og vi er altid åbne
							for nye samarbejder – særligt med koncertsteder og kulturaktører, der deler vores
							vision om at styrke unges adgang til musik og fællesskab. Vi søger kommercielle
							partnerskaber med venues og kulturinstitutioner på hele Fyn, og vi er interesserede i
							samarbejdspartnere, der kan bidrage med viden, rammer, netværk eller ressourcer
							inden for vores mærkesager.
							<br />
							<br />
							Vil du være med til at støtte vækstlaget og engagere flere unge i kulturlivet?
							Så <a href="mailto:konnekt.samarbejde@gmail.com">send os en mail</a>.
							Vi tager gerne en uforpligtende snak.
						</Accordion>
					</FAQSection>
				</section>
			</main >
		</>
	)
}

const Section = ({ children }: PropsWithChildren) => (
	<section className="prose-lg prose-headings:font-heading prose-headings:font-bold prose-p:text-text/75">
		{children}
	</section>
)

type FAQSectionProps = {
	title: string;
}

const FAQSection = ({ title, children }: PropsWithChildren<FAQSectionProps>) => (
	<section>
		<h2 className="font-semibold text-2xl mb-8">{title}</h2>
		<div className="flex flex-col gap-4">
			{children}
		</div>
	</section>
)
