import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { EditEventForm } from "~/components/events/edit-event-form/edit-event-form";
import { fetchEventByID } from "~/lib/event/event";
import { fetchAllGenres } from "~/lib/genre/genre";

export const loader = async ({ request }: LoaderFunctionArgs) => {
  const url = new URL(request.url)

  const id = url.searchParams.get("id") || ""

  const event = await fetchEventByID(parseInt(id), { headers: request.headers })
  const genres = await fetchAllGenres({ headers: request.headers })

  return { event, genres }
}

const EditEventPage = () => {
  const { event, genres } = useLoaderData<typeof loader>()
  console.log(genres)

  return (
    <main className="min-h-sub-nav px-auto py-16">
      <EditEventForm event={event} genres={genres} className="max-w-xl" />
    </main>
  )
}

export default EditEventPage;
