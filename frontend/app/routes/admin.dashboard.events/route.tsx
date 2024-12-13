import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { useToast } from "@/lib/context/toast.provider";
import { BiPlus } from "react-icons/bi";
import { EventEntry } from "@/routes/admin.dashboard.events/event-entry";
import { fetchEvents } from "@/lib/event";
import { useQuery } from "@tanstack/react-query";
import { startOfToday } from "date-fns";

const EventsPage = () => {
  const { addToast } = useToast()

  const { isPending, error, data: eventsResult, refetch } = useQuery({
    queryKey: ["events"],
    queryFn: () => fetchEvents({ limit: 8, fromDate: startOfToday() })
  })

  const deleteEvent = async (id: number) => {
    if (!confirm(`Delete event with id ${id}?`)) {
      return
    }

    const res = await fetch(`${window.ENV.BACKEND_URL}/events/${id}`, {
      method: "DELETE",
      credentials: "include"
    })

    if (res.ok) {
      addToast("Event slettet.", "success")
    } else {
      addToast("Kunne ikke slette event.", "error")
    }

    refetch()
  }

  if (error) return "Error"

  return (
    <Card className="min-w-xl">
      <CardHeader>
        <div className="flex items-center justify-between">
          <h2 className="font-bold text-xl">Kommende events.</h2>
          <a href="/admin/events/edit"><BiPlus className="text-xl" /></a>
        </div>
      </CardHeader>
      <CardContent>
        {isPending ? (<p className="italic text-foreground/50">Loader events...</p>) : (
          eventsResult.events.length > 0 ? (
            <div className="relative overflow-y-scroll">
              {eventsResult.events.map(event => (
                <EventEntry key={event.id} event={event} onDelete={() => deleteEvent(event.id)} />
              ))}
            </div>
          ) : (
            <p className="italic text-foreground/50">Ingen kommende events...</p>
          ))}
      </CardContent>
    </Card>
  )
}

export default EventsPage;
