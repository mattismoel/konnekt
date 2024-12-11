import { useLoaderData, useRevalidator } from "@remix-run/react";
import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { eventSchema } from "@/lib/dto/event.dto";
import { useToast } from "@/lib/context/toast.provider";
import env from "@/config/env";
import { BiPlus } from "react-icons/bi";
import { EventEntry } from "@/routes/admin.dashboard.events/event-entry";

export const loader = async () => {
  const res = await fetch(`${env.BACKEND_URL}/events`)

  if (!res.ok) {
    throw new Error(`Could not fetch events: ${res.status}`)
  }

  const data = await res.json()

  const events = eventSchema.array().parse(data)

  return {
    events
  }
}

const EventsPage = () => {
  const { addToast } = useToast()
  const { events } = useLoaderData<typeof loader>()
  const { revalidate } = useRevalidator()

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

    revalidate()
  }

  return (
    <Card>
      <CardHeader>
        <div className="flex items-center justify-between">
          <h2 className="font-bold text-xl">Kommende events.</h2>
          <a href="/admin/events/edit"><BiPlus className="text-xl" /></a>
        </div>
      </CardHeader>
      <CardContent>
        {events.length > 0 ? (
          <div className="relative overflow-y-scroll">
            {events.map(event => (
              <EventEntry key={event.id} event={event} onDelete={() => deleteEvent(event.id)} />
            ))}
          </div>
        ) : (
          <p className="italic text-foreground/50">Ingen kommende events...</p>
        )}
        {/*
          <Fader
            direction="to-top"
            color="neutral-950"
            className="h-20 absolute bottom-0 left-0"
          />
        */}
      </CardContent>
    </Card>
  )
}

export default EventsPage;
