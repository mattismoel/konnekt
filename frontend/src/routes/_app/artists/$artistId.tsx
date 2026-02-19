import Fader from "@/lib/components/fader";
import PageMeta from "@/lib/components/page-meta";
import SpotifyPreview from "@/lib/components/spotify-preview";
import { socialUrlToIcon } from "@/lib/features/artist/artist";
import {
  createArtistByIdOpts,
  createArtistEventsOpts,
} from "@/lib/features/artist/query";
import EventGrid from "@/lib/features/event/components/event-grid";
import { trackIdFromUrl } from "@/lib/spotify";
import { useSuspenseQuery } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_app/artists/$artistId")({
  component: RouteComponent,
  loader: async ({
    context: { queryClient },
    params: { artistId: artistIdStr },
  }) => {
    const artistId = parseInt(artistIdStr);
    const artistOptions = createArtistByIdOpts(artistId);
    const artistEventsOpts = createArtistEventsOpts(artistId);

    queryClient.ensureQueryData(artistEventsOpts);
    queryClient.ensureQueryData(artistOptions);

    return { artistOptions, artistEventsOpts };
  },
});

function RouteComponent() {
  const { artistOptions, artistEventsOpts } = Route.useLoaderData();

  const { data: artist } = useSuspenseQuery(artistOptions);
  const {
    data: { records: artistEvents },
  } = useSuspenseQuery(artistEventsOpts);

  if (!artist) return <p>No such artist...</p>;

  let trackId = artist.previewUrl
    ? trackIdFromUrl(artist.previewUrl)
    : undefined;
  return (
    <>
      <PageMeta
        title={`Konnekt | ${artist.name}`}
        description={`Oplev ${artist.name} til "${artistEvents.at(0)?.title}"`}
      />
      <main>
        <div className="min-h-svh">
          <div className="relative isolate flex min-h-[85svh] items-end py-16">
            <img
              src={artist.imageUrl}
              alt="Cover af {artist.name}"
              className="absolute top-0 left-0 h-full w-full object-cover"
            />
            <Fader
              direction="right"
              className="absolute hidden w-96 from-black md:block"
            />
            <Fader direction="up" className="absolute h-[512px] from-black" />

            <div className="z-10 mx-responsive flex w-full flex-col items-start justify-between gap-8 md:flex-row md:items-end">
              <h1
                style={{ wordSpacing: "100vw" }}
                className="font-heading text-7xl font-bold text-heading text-shadow-md md:text-8xl lg:text-9xl"
              >
                {artist.name}
              </h1>
              <div className="flex gap-4 text-3xl text-shadow-sm">
                {artist.socials.map((social) => {
                  const Icon = socialUrlToIcon(social);
                  return (
                    <a
                      key={social}
                      href={social}
                      className="transition-colors hover:text-text-light"
                    >
                      <Icon />
                    </a>
                  );
                })}
              </div>
            </div>
          </div>

          <div className="border-t border-t-zinc-900">
            <article className="mx-responsive w-full space-y-16 bg-zinc-950 py-16">
              <section className="space-y-8">
                <div
                  className="prose prose-lg max-w-none prose-invert md:prose-base"
                  dangerouslySetInnerHTML={{ __html: artist.description }}
                />
                {trackId && <SpotifyPreview trackId={trackId} />}
              </section>

              {artistEvents.length > 0 && (
                <section>
                  <h1 className="mb-8 font-heading text-2xl font-bold text-heading">
                    Oplev {artist.name} her
                  </h1>
                  <EventGrid events={artistEvents} />
                </section>
              )}
            </article>
          </div>
        </div>
      </main>
    </>
  );
}
