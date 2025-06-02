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
            Konnekt er et projekt med henblik på unge musikere.
            Det danske musikmiljø er for svært at bryde igennem - især for unge
            aspirerende musikere og det skal ændres.
          </p>
        </Section>

        <Section>
          <h1>Missionen</h1>
          <p className="leading-loose">
            Konnekt giver unge musikere og upcoming bands en scene og et
            publikum gennem koncertarrangementer. Hver koncert har to hovednavne:
            én ny og én mere etableret upcoming kunstner.
            Derudover optræder unge, lokale talenter som opvarmning.

            Formålet er at engagere flere unge i Odenses kulturliv ved at
            præsentere ny musik inden for genrer, de allerede interesserer sig
            for, samt skabe rum for inspiration, rådgivning og udvikling gennem
            talks.
          </p>
        </Section>

        <section className="flex flex-col">
          <h1 className="text-center font-heading text-4xl font-bold mb-16">Mød holdet</h1>
          <TeamDisplay allTeams={teams} members={members} />
        </section>

        <section>
          <h1 className="font-heading text-4xl font-bold mb-8">Ofte stillede spørgsmål</h1>
          <div className="flex flex-col gap-4">
            <Accordion title="Hvor ofte afholder Konnekt events?">
              Vi bestræber os på, at afholde flest mulige events af højest mulig
              kvalitet. Vores events består af de dygtigste upcoming kunstnere, som
              vores event management arbejder på højtryk for at finde.
            </Accordion>
            <Accordion title="Jeg er kunstner. Hvordan bliver jeg booket?">
              Hvis du er musiker, og brænder for at komme på scenen, så tag fat i
              os på vores booking mail <a href="mailto:booking.konnekt@gmail.com">booking.konnekt@gmail.com</a>
            </Accordion>
            <Accordion title="Hvordan kan jeg hjælpe foreningen?">
              Du er altid velkommen til at kontakte os på <a href="mailto:konnekt.samarbejde@gmail.com">
                konnekt.samarbejde@gmail.com
              </a>, for at blive medlem af foreningen. Vi kan garanteret bruge din
              hjælp!
            </Accordion>
            <Accordion title="Hvad med kunstnere fra Sjælland og Jylland?">
              Vi beskæftiger os som udgangspunkt med fysnke kunstnere, men ser også
              gerne samarbejde med ikke-fynske kunstnere. Tag fat i os på&nbsp;<a href="mailto:booking.konnekt@gmail.com">booking.konnekt@gmail.com</a>,
              hvis du er interesseret i at stå på scenen.
            </Accordion>
          </div>
        </section>
      </main>
    </>
  )
}

const Section = ({ children }: PropsWithChildren) => (
  <section className="prose-base prose-headings:font-heading prose-headings:font-bold prose-p:text-text/75">
    {children}
  </section>

)
