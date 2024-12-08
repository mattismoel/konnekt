import { useLoaderData } from "@remix-run/react"
import { EventGrid } from "~/components/events/event-grid"
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
    <main className="px-auto py-20 h-sub-nav">
      <h1 className="text-2xl font-bold mb-4">Kommende events.</h1>
      <EventGrid events={events} />
    </main>
  )
}

export default EventsPage
