import { useParams } from "@remix-run/react"
import { EventDTO } from "@/lib/dto/event.dto"

import { EventCalendar } from "@/components/events/event-calendar/event-calendar";
import { EventDetails } from "./event-details";
import { EventCaroussel } from "@/components/events/event-caroussel";
import { useEffect, useState } from "react";
import { fetchEventByID, fetchEvents } from "@/lib/event";
import { endOfWeek, startOfDay, startOfWeek } from "date-fns";

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

    const { events: weeklyEvents } = await fetchEvents({
      fromDate: startOfWeek(event.fromDate),
      toDate: endOfWeek(event.fromDate)
    })

    setWeeklyEvents(weeklyEvents)

    const { events: recommendedEvents } = await fetchEvents({
      limit: 5,
      fromDate: startOfDay(new Date()),
    })

    // Filter out the actual event from the recommended.
    setRecommendedEvents(recommendedEvents.filter(e => e.id !== event.id))

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
        <EventCaroussel events={recommendedEvents} />
      </div>
    </main>
  )
}

export default EventPage;
