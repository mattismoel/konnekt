import env from "~/config/env";
import { useLoaderData } from "@remix-run/react";
import { Card, CardContent, CardHeader } from "~/components/ui/card"
import { eventSchema } from "~/lib/event/event.dto";
import { EventEntry } from "./event-entry";

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
  const { events } = useLoaderData<typeof loader>()

  const deleteEvent = (id: number) => {
    confirm(`Delete event with id ${id}?`)
  }

  return (
    <Card>
      <CardHeader>
        <h1>Events.</h1>
      </CardHeader>
      <CardContent>
        <h2 className="font-bold text-xl mb-4">Kommende events.</h2>
        <div className="relative overflow-y-scroll">
          {events.map(event => (
            <EventEntry event={event} onDelete={deleteEvent} />
          ))}
        </div>
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
