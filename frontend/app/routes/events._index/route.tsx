import { EventGrid } from "@/components/events/event-grid"
import { fetchEvents } from "@/lib/event"
import { Paginator } from "@/components/ui/paginator"
import { useSearchParams } from "@remix-run/react"
import { useQuery } from "@tanstack/react-query"

const PAGE_SIZE = 6

const EventsPage = () => {
  const [searchParams, setSearchParams] = useSearchParams()
  const page = parseInt(searchParams.get("page") || "1", 10)


  const { isPending, error, data } = useQuery({
    queryKey: ["events"],
    queryFn: () => fetchEvents({
      page,
      pageSize: PAGE_SIZE
    }),
  })

  if (error) return "Error"


  return (
    <main className="px-auto py-20 min-h-sub-nav">
      <h1 className="text-2xl font-bold mb-4">Kommende events.</h1>
      <EventGrid className="mb-4" events={data?.events || []} isLoading={isPending} />
      <Paginator
        totalPages={Math.ceil((data?.totalSize || 1) / PAGE_SIZE) || 1}
        currentPage={page}
        onSelect={(newPage) => setSearchParams({ page: newPage.toString() })}
      />
    </main>
  )
}

export default EventsPage
