import env from "@/config/env"
import { LoaderFunctionArgs } from "@remix-run/node"
import { useLoaderData } from "@remix-run/react"
import { EventDTO, eventSchema } from "@/lib/dto/event.dto"

import { EventCalendar } from "@/components/events/event-calendar/event-calendar";
import { EventDetails } from "./event-details";
import { EventCaroussel } from "@/components/events/event-caroussel";

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
  const { event, weeklyEvents, recommendedEvents } = useLoaderData<typeof loader>()
  const { title, description, coverImageUrl } = event;

  return (
    <main className="min-h-sub-nav pb-16 bg-black text-white">
      <div className="relative h-[calc((100vh/4)*3)] overflow-hidden">
        <img
          src={coverImageUrl}
          alt={title}
          className="h-full w-full object-cover"
        />
        <div
          className="absolute bottom-0 left-0 h-2/3 w-full bg-gradient-to-t from-black"
        ></div>
        <div className="absolute w-full bottom-0 left-0 p-8 flex flex-col">
          <h1 className="w-full text-5xl md:text-8xl font-bold mb-2">{event.title}</h1>
          <EventDetails event={event} />
        </div>
      </div>

      <article dangerouslySetInnerHTML={{ __html: description }} className="px-auto pt-20 pb-16 text-gray-400 prose prose-invert max-w-none">
      </article>
      <div className="px-auto">
        <h2 className="font-bold text-2xl">Billetten gælder også til.</h2>
        <EventCalendar activeID={event.id} events={weeklyEvents} className="h-64" />
        <h1 className="font-bold text-4xl mb-12">Se også.</h1>
        <EventCaroussel events={[]} />
      </div>
    </main>
  )
}

export default EventPage;
