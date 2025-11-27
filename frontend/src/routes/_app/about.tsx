import PageMeta from '@/lib/components/page-meta'
import TeamDisplay from '@/lib/components/team-display'
import { membersQueryOpts, teamsQueryOpts } from '@/lib/features/auth/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

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
			<main className='min-h-svh py-32 px-auto flex flex-col'>
				<h1 className="font-heading text-4xl font-bold mb-4">Hvem er vi?</h1>
				<p className="leading-relaxed text-text/75 mb-16">
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

				<section className="flex flex-col">
					<h1 className="text-center font-heading text-4xl font-bold mb-16">Mød holdet</h1>
					<TeamDisplay allTeams={teams} members={members} />
				</section>
			</main>
		</>
	)
}
