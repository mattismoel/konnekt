import { createFileRoute } from '@tanstack/react-router'

import LandingVideo from '@/lib/assets/landing.mp4';
import Caroussel from '@/lib/components/caroussel';
import EventCard from '@/lib/components/event-card';
import LogoScroller from '@/lib/components/logo-scroller';

import OdenseKommuneLogo from '@/lib/assets/logos/odense-kommune-logo.svg';
import UngOdenseLogo from '@/lib/assets/logos/ungodense-logo.svg';
import PostenLogo from '@/lib/assets/logos/posten-logo.svg';
import KulturMaskinenLogo from '@/lib/assets/logos/kulturmaskinen-logo.svg';
import SpillestedetOdenseLogo from '@/lib/assets/logos/spillestedet-odense-logo.svg';

import Fader from '@/lib/components/fader';
import GlowCursor from '@/lib/components/glow-cursor';
import { FaArrowRight } from 'react-icons/fa';

import { useListUpcomingEvents } from '@/lib/features/hook';
import LinkButton from '@/lib/components/ui/button/link-button';


export const Route = createFileRoute('/_app/')({
  component: App,
})

function App() {
  const { data, isLoading } = useListUpcomingEvents()

  if (isLoading) return <p>...</p>

  return (
    <div>
      <section className="px-auto -z-50 flex h-svh flex-col justify-center gap-16">
        <div
          className="pointer-events-none absolute top-0 left-0 isolate -z-10 h-full w-full overflow-hidden"
        >
          <GlowCursor />
          <video
            loop
            muted
            autoPlay
            src={LandingVideo}
            className="pointer-events-none fixed top-0 left-0 z-0 h-full w-full object-cover brightness-50"
          >
            <track kind="captions" />
          </video>
        </div>
        <section className="flex max-w-lg flex-col gap-8">
          <h2 className="font-heading text-5xl">For et stærkere <b>fynsk musisk vækstlag</b></h2>
          <p className="text-text/75">
            En forening med formål, at støtte det lokale fynske musiske vækstlag og give aspirerende
            musikere et springbræt til den danske musikscene.
          </p>
          <div className="flex w-full flex-col-reverse gap-4 sm:flex-row">
            <LinkButton to="/about" variant="outline" className="w-full sm:w-fit">Læs mere</LinkButton>
            <LinkButton to="/events" className="group w-full items-center gap-2 sm:w-fit">
              Se events
              <FaArrowRight className="text-sm transition-transform group-hover:translate-x-1" />
            </LinkButton>
          </div>
        </section>
        <Fader direction="up" className="absolute h-64 from-black/75" />
      </section>

      <section className="bg-zinc-950">
        {data?.records && data.records.length > 0 && (
          <section className="px-auto py-16">
            <h1 className="font-heading mb-8 text-2xl font-bold">Kommende events</h1>
            <Caroussel>
              {data?.records.map(event => (
                <EventCard key={event.id} event={event} />
              ))}
            </Caroussel>
          </section>

        )}

        <section className="px-auto space-y-8 py-16">
          {/*  MISSION STATEMENT  */}
          <section>
            <h1 className="font-heading mb-8 text-2xl font-bold">Vores mission</h1>
            <p className="text-text/75">
              Foreningen Konnekt har som formål at støtte unge musikere og skabe en platform, hvor de
              kan vise deres talent frem og få vigtig erfaring med liveoptrædener. Projektet skal gøre
              det nemmere for spirende talenter at finde deres plads i musikmiljøet og opbygge et
              publikum.
              <br />
              <br />
              Samtidig ønsker vi at give publikum – især unge – mulighed for at opdage nye kunstnere i genrer,
              de allerede har stiftet bekendtskab med før. Derudover vil vi styrke musikmiljøet i Odense
              ved at skabe et fællesskab mellem nye og mere erfarne upcoming kunstnere, som kan dele erfaringer,
              inspirere hinanden og måske endda finde samarbejdspartnere. På den måde skaber Konnekt ikke
              kun koncertoplevelser, men også en grobund for kreativ udvikling og vækst i den lokale kultur.
            </p>
          </section>

          {/*  SPONSORS  */}
          <section className="z-0 flex w-full flex-col gap-8">
            <span className="font-bold">Med støtte fra</span>
            <div className="relative isolate w-full">
              <Fader direction="right" className="absolute z-50 w-32 from-zinc-950" />
              <Fader direction="left" className="absolute z-50 w-32 from-zinc-950" />
              <LogoScroller
                className="h-10 w-full"
                srcs={new Map<string, string>([
                  ['Spillestedet Odense', SpillestedetOdenseLogo],
                  ['UngOdense', UngOdenseLogo],
                  ['Posten', PostenLogo],
                  ['Kulturmaskinen', KulturMaskinenLogo],
                  ['Odense Kommune', OdenseKommuneLogo]
                ])}
              />
            </div>
          </section>
        </section>
      </section>
    </div >
  )
}
