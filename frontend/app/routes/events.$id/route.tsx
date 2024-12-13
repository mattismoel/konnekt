import { useParams } from "@remix-run/react"
import { EventDTO } from "@/lib/dto/event.dto"

import { EventCalendar } from "@/components/events/event-calendar/event-calendar";
import { EventDetails } from "./event-details";
import { EventCaroussel } from "@/components/events/event-caroussel";
import { useEffect, useState } from "react";
import { fetchEventByID, fetchEvents } from "@/lib/event";
import { endOfWeek, startOfDay, startOfToday, startOfWeek } from "date-fns";
import { useQueries, useQuery } from "@tanstack/react-query";

const EventPage = () => {
  const { id } = useParams()

  const { data: event } = useQuery({
    queryKey: ["events", id],
    queryFn: () => fetchEventByID(parseInt(id || "0"))
  })

  const { data: recommendedEventsResult } = useQuery({
    queryKey: ["events"],
    queryFn: () => fetchEvents({
      fromDate: startOfWeek(event?.fromDate || new Date()),
      toDate: endOfWeek(event?.fromDate || new Date())
    }),
    enabled: !!event
  })

  const { isPending, error, data: weeklyEventsResult } = useQuery({
    queryKey: ["events"],
    queryFn: () => fetchEvents({
      limit: 5,
      fromDate: startOfToday()
    }),
    enabled: !!event
  })

  if (error) return "Error"

  return (
    <main className="min-h-sub-nav pb-16 bg-black text-white">
      <EventDetails event={event || undefined} isLoading={isPending} />
      <article dangerouslySetInnerHTML={{ __html: event?.description || "" }} className="px-auto pt-20 pb-16 text-gray-400 prose prose-invert max-w-none">
      </article>
      <div className="px-auto">
        <h2 className="font-bold text-2xl">Billetten gælder også til.</h2>
        <EventCalendar activeID={event?.id} events={weeklyEventsResult?.events || []} className="h-64" />
        <h1 className="font-bold text-4xl mb-12">Se også.</h1>
        <EventCaroussel events={recommendedEventsResult?.events || []} />
      </div>
    </main>
  )
}

export default EventPage;
