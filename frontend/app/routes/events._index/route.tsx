import { useLoaderData } from "@remix-run/react"
import { EventCard } from "~/components/events/event-card"
import env from "~/config/env"
import { eventSchema } from "~/lib/event/event.dto"

export const loader = async () => {
  try {
    const res = await fetch(`${env.BACKEND_URL}/events`)

    if (!res.ok) {
      throw new Error(`Error fetching backend: ${res.status}, ${res.statusText}`)
    }

    const data = await res.json()

    const events = eventSchema.array().parse(data)

    return events
  } catch (e) {
    console.error(e)
    return []
  }
}

const EventsPage = () => {
  const events = useLoaderData<typeof loader>();

  return (
    <main>
      {events.map(event => (<EventCard key={event.id} event={event} />))}
    </main>
  )
}

export default EventsPage
