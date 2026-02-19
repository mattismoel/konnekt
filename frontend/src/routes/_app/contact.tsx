import Accordion from "@/lib/components/accordion";
import InlineLink from "@/lib/components/inline-link";
import Switch from "@/lib/components/switch";
import { createFileRoute } from "@tanstack/react-router";
import { useMemo, useState } from "react";

const categoryNames = ["booking", "volunteers", "partners", "other"] as const;
type CategoryName = (typeof categoryNames)[number];

const categoryDetailsMap = new Map<CategoryName, { email: string }>([
  ["booking", { email: "booking.konnekt@gmail.com" }],
  ["volunteers", { email: "konnekt.samarbejde@gmail.com" }],
  ["partners", { email: "konnekt.samarbejde@gmail.com" }],
  ["other", { email: "konnekt.samarbejde@gmail.com" }],
]);

export const Route = createFileRoute("/_app/contact")({
  component: RouteComponent,
});

function RouteComponent() {
  const [selectedCategory, setSelectedCategory] =
    useState<CategoryName>("booking");

  const categoryDetails = useMemo(
    () => categoryDetailsMap.get(selectedCategory),
    [selectedCategory],
  );

  const handleSwitchCategory = (id: string) => {
    const categoryName = categoryNames.find((category) => category === id);
    if (!categoryName) throw new Error("Not a valid category name");
    setSelectedCategory(categoryName);
  };

  return (
    <main className="mx-responsive flex min-h-svh flex-col gap-32 py-32">
      <section>
        <h1 className="mb-4 font-heading text-4xl font-bold text-text-light">
          Er du i tvivl om noget?
        </h1>
        <p className="text-text/75 mb-16 leading-relaxed">
          Lad os se, om ikke vi kan besvare dine spørgsmål. Hvis ikke dit
          spørgsmål kan findes herunder, så skal du være velkommen til at tage
          fat i os&nbsp;&mdash;&nbsp;så svarer vi hurtigst muligt.
        </p>

        <Switch
          entries={[
            { id: "booking", value: "Booking" },
            { id: "volunteers", value: "Frivillige" },
            { id: "partners", value: "Samarbejdspartnere" },
          ]}
          prefix="Jeg vil vide noget om..."
          selected={selectedCategory}
          onSwitch={handleSwitchCategory}
          className="mb-8 w-full"
        />

        <div className="mb-8 flex min-h-42 flex-col gap-2">
          {selectedCategory === "booking" ? (
            <>
              <Accordion title="Hvordan bliver jeg booket?">
                Konnekt er altid på udkig efter nye, spændende navne til vores
                koncerter. Hvis du eller dit band har lyst til at spille, kan I{" "}
                <InlineLink href="mailto:booking.konnekt@gmail.com">
                  sende en mail
                </InlineLink>{" "}
                med lidt info om jer selv, og et lydklip fra en øver, koncert
                eller demo til. Vi gennemgår løbende alle henvendelser og tager
                jer med i overvejelserne til kommende koncerter.
              </Accordion>
              <Accordion title="Kan man blive booket, hvis man ikke er fra Fyn?">
                Ja, det kan du godt! Konnekt har som kerneværdi at styrke det
                fynske musik vækstlag, men tager også gerne imod bud fra resten
                af landet.
              </Accordion>
              <Accordion title="Kan man blive booket, hvis man ikke har udgivet musik?">
                Ja, det kan man godt! Konnekt booker udelukkende bands/musikere
                der optræder med selvskrevet musik - også selvom du/i ikke har
                udgivet musik.
              </Accordion>
            </>
          ) : selectedCategory === "volunteers" ? (
            <>
              <Accordion title="Hvordan kan jeg hjælpe?">
                Konnekt er drevet af unge kræfter og gode idéer, og vi er altid
                på udkig efter flere. Uanset om du vil arrangere, planlægge,
                markedsføre, økonomien, skabe visuals, booke bands, skabe en god
                atmosfære, eller noget helt andet, så er der plads til dig.
                <br />
                <br />
                Vi tror på, at kulturlivet i Odense bliver stærkere, når flere
                er med til at forme det. Derfor lytter vi til nye perspektiver
                og er altid åbne for friske inputs og kreative hjerner. Har du
                lyst til at være med, så{" "}
                <InlineLink href="mailto:konnekt.samarbejde@gmail.com">
                  skriv til os
                </InlineLink>
                , og fortæl hvad du brænder for. Vi glæder os til at høre fra
                dig.
              </Accordion>
              <Accordion title="Hvad laver man som frivillig?">
                Som frivillig i Konnekts bestyrelse er du med til at forme alt
                fra idé til virkelighed. Du deltager i møder, planlægger
                koncerter og udvikler på, hvordan vi bedst løfter ungt kulturliv
                i Odense. Det kan være booking, PR, økonomi, samarbejder eller
                hvordan vi gør publikum gladere – alt efter hvad du brænder for.
                <br />
                <br />
                Når koncertdagen kommer, hjælper du med det praktiske:
                kunstnerforplejning, opsætning, afvikling og alt det, der får
                aftenen til at spille. Du får erfaring, indflydelse og et
                fællesskab med andre unge, der ligesom dig gerne vil gøre en
                forskel.
                <br />
                <br />
                Har du lyst til at være med? Så{" "}
                <InlineLink href="mailto:konnekt.samarbejde@gmail.com">
                  send os endelig en mail
                </InlineLink>
                . Vi vil gerne høre fra dig!
              </Accordion>
            </>
          ) : (
            <>
              <Accordion title="Hvordan kan vi samarbejde?">
                Konnekt er bygget på et grundlag af stærke samarbejdspartnere,
                og vi er altid åbne for nye samarbejder – særligt med
                koncertsteder og kulturaktører, der deler vores vision om at
                styrke unges adgang til musik og fællesskab. Vi søger
                kommercielle partnerskaber med venues og kulturinstitutioner på
                hele Fyn, og vi er interesserede i samarbejdspartnere, der kan
                bidrage med viden, rammer, netværk eller ressourcer inden for
                vores mærkesager.
                <br />
                <br />
                Vil du være med til at støtte vækstlaget og engagere flere unge
                i kulturlivet? Så{" "}
                <InlineLink href="mailto:konnekt.samarbejde@gmail.com">
                  send os en mail
                </InlineLink>
                . Vi tager gerne en uforpligtende snak.
              </Accordion>
            </>
          )}
        </div>

        <div className="w-full">
          <p className="text-text/75 leading-relaxed">
            Fik du ikke svar? Kontakt os på&nbsp;
            <InlineLink href={`mailto:${categoryDetails?.email}`}>
              {categoryDetails?.email}
            </InlineLink>
            .
          </p>
        </div>
      </section>
    </main>
  );
}
