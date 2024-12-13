import env from "@/config/env"
import { LoaderFunctionArgs } from "@remix-run/node"
import { useParams } from "@remix-run/react"
import { EventDTO, eventSchema } from "@/lib/dto/event.dto"

import { EventCalendar } from "@/components/events/event-calendar/event-calendar";
import { EventDetails } from "./event-details";
import { EventCaroussel } from "@/components/events/event-caroussel";
import { useEffect, useState } from "react";
import { fetchEventByID } from "@/lib/event";

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const id = parseInt(params.id || "")

  const res = await fetch(`${env.BACKEND_URL}/events/${id}`)

  if (!res.ok) {
    throw new Error(`Could not fetch for event with id ${id}: ${res.status}`)
  }

  const data = await res.json()

  const event = eventSchema.parse(data)
  return {
    event,
    weeklyEvents: [] as EventDTO[],
    recommendedEvents: [] as EventDTO[],
  }
}

const EventPage = () => {
  const { id } = useParams()
  const [event, setEvent] = useState<EventDTO | undefined>(undefined)
  const [weeklyEvents, setWeeklyEvents] = useState<EventDTO[]>([])
  const [recommendedEvents, setRecommendedEvents] = useState<EventDTO[]>([])

  const [loading, setLoading] = useState(true)

  const handleFetchEvent = async () => {
    setLoading(true)
    const event = await fetchEventByID(parseInt(id || "1"))

    if (!event) {
      throw new Error("Could not load event")
    }

    setEvent(event)
    setLoading(false)
  }

  useEffect(() => {
    handleFetchEvent()
  }, [])

  return (
    <main className="min-h-sub-nav pb-16 bg-black text-white">
      <EventDetails event={event} isLoading={loading} />
      <article dangerouslySetInnerHTML={{ __html: event?.description || "" }} className="px-auto pt-20 pb-16 text-gray-400 prose prose-invert max-w-none">
      </article>
      <div className="px-auto">
        <h2 className="font-bold text-2xl">Billetten gælder også til.</h2>
        <EventCalendar activeID={event?.id} events={weeklyEvents} className="h-64" />
        <h1 className="font-bold text-4xl mb-12">Se også.</h1>
        <EventCaroussel events={[]} />
      </div>
    </main>
  )
}

export default EventPage;
