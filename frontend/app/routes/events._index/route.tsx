import { EventGrid } from "@/components/events/event-grid"
import { EventDTO } from "@/lib/dto/event.dto"
import { useEffect, useState } from "react"
import { fetchEvents } from "@/lib/event"

const PAGE_SIZE = 12

const EventsPage = () => {
  const [events, setEvents] = useState<EventDTO[]>([])
  const [pageCount, setPageCount] = useState(0)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const handleFetchEvents = async () => {
      setLoading(true)
      const { events, totalSize } = await fetchEvents()

      setTimeout(() => {
        setEvents(events)
        setPageCount(totalSize % PAGE_SIZE)
        setLoading(false)
      }, 2000)
    }

    handleFetchEvents()
  }, [])

  return (
    <main className="px-auto py-20 h-sub-nav">
      <h1 className="text-2xl font-bold mb-4">Kommende events.</h1>
      <EventGrid events={events} loading={loading} />
      {pageCount}
    </main>
  )
}

export default EventsPage
