import { CreateEditEventDTO, EventDTO } from "@/lib/dto/event.dto";
import { useToast } from "@/lib/context/toast.provider";
import { EditEventForm } from "./edit-event-form";
import { useEffect, useState } from "react";
import { fetchEventByID } from "@/lib/event";
import { useSearchParams } from "@remix-run/react";
import { fetchAllGenres } from "@/lib/genre";
import { LoadingScreen } from "@/components/ui/loading-screen";
import { sleep } from "@/lib/time";

const EditEventPage = () => {
  const { addToast } = useToast()
  //const { event, genres } = useLoaderData<typeof loader>()

  const [searchParams] = useSearchParams()

  const [event, setEvent] = useState<EventDTO | undefined>(undefined)
  const [genres, setGenres] = useState<string[]>([])
  const [loading, setLoading] = useState(false)


  const handleFetch = async () => {
    setLoading(true)

    const genres = await fetchAllGenres()
    setGenres(genres)

    const id = searchParams.get("id")

    if (!id) {
      setLoading(false)
      return
    }

    const event = await fetchEventByID(parseInt(id))

    if (!event) {
      throw new Error("No event found")
    }

    setEvent(event)
    setLoading(false)
  }

  useEffect(() => {
    handleFetch()
  }, [])

  const handleSubmit = async (data: CreateEditEventDTO) => {
    let res = await fetch(`${window.ENV.BACKEND_URL}/events`, {
      method: "POST",
      credentials: "include",
      body: JSON.stringify(data),
      headers: { "Content-Type": "application/json" }
    })

    if (!res.ok) {
      addToast("Kunne ikke uploade event", "error")
      throw new Error(`Could not upload base event: ${res.statusText}`)
    }

    addToast("Event uploadet", "success")
  }

  if (loading) {
    return <LoadingScreen label="Loader event..." />
  }

  return (
    <main className="min-h-sub-nav px-auto py-16">
      <EditEventForm
        event={event || null}
        genres={genres}
        className="max-w-xl"
        onSubmit={handleSubmit}
      />
    </main>
  )
}

export default EditEventPage;
