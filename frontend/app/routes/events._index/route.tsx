import { EventGrid } from "@/components/events/event-grid"
import { EventDTO } from "@/lib/dto/event.dto"
import { useEffect, useState } from "react"
import { fetchEvents } from "@/lib/event"
import { Paginator } from "@/components/ui/paginator"
import { useSearchParams } from "@remix-run/react"

const PAGE_SIZE = 6

const EventsPage = () => {
  const [searchParams, setSearchParams] = useSearchParams()
  const page = parseInt(searchParams.get("page") || "1", 10)

  const [events, setEvents] = useState<EventDTO[]>([])
  const [totalPageCount, setTotalPageCount] = useState(0)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const handleFetchEvents = async () => {
      setLoading(true)
      const { events, totalSize } = await fetchEvents({
        page,
        pageSize: PAGE_SIZE,
      })

      setEvents(events)
      let pageCount = Math.floor(totalSize / PAGE_SIZE)
      if (pageCount === 0) pageCount = 1
      setTotalPageCount(pageCount)
      setLoading(false)
    }

    handleFetchEvents()
  }, [page])

  return (
    <main className="px-auto py-20 min-h-sub-nav">
      <h1 className="text-2xl font-bold mb-4">Kommende events.</h1>
      <EventGrid className="mb-4" events={events} loading={loading} />
      <Paginator
        totalPages={totalPageCount}
        currentPage={page}
        onSelect={(newPage) => setSearchParams({ page: newPage.toString() })}
      />
    </main>
  )
}

export default EventsPage
