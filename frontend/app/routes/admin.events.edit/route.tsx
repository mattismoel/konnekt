import { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { fetchEventByID } from "@/lib/event";
import { CreateEditEventDTO } from "@/lib/dto/event.dto";
import { fetchAllGenres } from "@/lib/genre";
import { useToast } from "@/lib/context/toast.provider";
import { EditEventForm } from "./edit-event-form";

export const loader = async ({ request }: LoaderFunctionArgs) => {
  const url = new URL(request.url)

  const id = url.searchParams.get("id") || ""

  const event = await fetchEventByID(parseInt(id), { headers: request.headers })
  const genres = await fetchAllGenres({ headers: request.headers })

  return { event, genres }
}

const EditEventPage = () => {
  const { addToast } = useToast()
  const { event, genres } = useLoaderData<typeof loader>()

  const handleSubmit = async (data: CreateEditEventDTO) => {
    console.log("Hello")
    let res = await fetch(`${window.ENV.BACKEND_URL}/events`, {
      method: "post",
      credentials: "include",
      body: JSON.stringify(data)
    })

    if (!res.ok) {
      addToast("Kunne ikke uploade event", "error")
      throw new Error(`Could not upload base event: ${res.statusText}`)
    }

    addToast("Event uploadet", "success")
  }

  return (
    <main className="min-h-sub-nav px-auto py-16">
      <EditEventForm
        event={event}
        genres={genres}
        className="max-w-xl"
        onSubmit={handleSubmit}
      />
    </main>
  )
}

export default EditEventPage;
