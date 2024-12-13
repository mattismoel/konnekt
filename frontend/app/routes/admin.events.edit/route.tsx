import { CreateEditEventDTO } from "@/lib/dto/event.dto";
import { useToast } from "@/lib/context/toast.provider";
import { EditEventForm } from "./edit-event-form";
import { fetchEventByID } from "@/lib/event";
import { useSearchParams } from "@remix-run/react";
import { fetchAllGenres } from "@/lib/genre";
import { LoadingScreen } from "@/components/ui/loading-screen";
import { useQuery } from "@tanstack/react-query";

const EditEventPage = () => {
  const { addToast } = useToast()

  const [searchParams] = useSearchParams()

  const id = parseInt(searchParams.get("id") || "0")

  const { data: event } = useQuery({
    queryKey: ["events", id],
    queryFn: () => fetchEventByID(id),
  })

  const { isPending, error, data: genres } = useQuery({
    queryKey: ["genres"],
    queryFn: () => fetchAllGenres(),
    enabled: !!event
  })

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

  if (error) return "Error"

  if (isPending) {
    return <LoadingScreen label="Loader event..." />
  }

  return (
    <main className="min-h-sub-nav px-auto py-16">
      <EditEventForm
        event={event || null}
        genres={genres || []}
        className="max-w-xl"
        onSubmit={handleSubmit}
      />
    </main>
  )
}

export default EditEventPage;
