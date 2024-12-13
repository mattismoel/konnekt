import { BiCalendar, BiMap, BiMusic } from "react-icons/bi"
import { EventDTO } from "@/lib/dto/event.dto"
import { formatDateString } from "@/lib/time"
import { cn } from "@/lib/utils";
import { Fader } from "@/components/ui/fader";

type Props = {
  isLoading: boolean;
  event: EventDTO | undefined
}

export const EventDetails = ({ event, isLoading }: Props) => {
  const { venue, fromDate } = event || {}

  if (isLoading) return <Skeleton />

  return (
    <div className="relative h-[calc((100vh/4)*3)] overflow-hidden">
      <img
        src={event?.coverImageUrl}
        alt={event?.title}
        className="absolute top-0 left-0 h-full w-full object-cover"
      />
      {/* FADER */}
      <Fader direction="to-bottom" />
      <div className="absolute bottom-0 left-0 pb-12 px-12 space-y-2">
        <h1 className={cn(
          "w-full text-5xl md:text-8xl font-bold mb-2",
        )}>{event?.title}</h1>
        <div className="flex gap-2 items-center">
          <BiMap />
          <span>{venue?.name}, {venue?.city}</span>
        </div>
        <div className="flex gap-2 items-center">
          <BiCalendar />
          <span>{formatDateString(fromDate || new Date())}</span>
        </div>
        <div className="flex gap-2 items-center">
          <BiMusic />
          <span>{event?.genres.join(", ")}</span>
        </div>
      </div>
    </div>
  )
}

const Skeleton = () => {
  return (
    <div className="relative h-[calc((100vh/4)*3)] overflow-hidden">
      {/* IMAGE */}
      <div className="absolute top-0 left-0 h-full w-full bg-background"></div>

      <Fader direction="to-bottom" />
      <div className="absolute bottom-0 left-0 px-12 pb-12 space-y-3 z-50">
        {/* TITLE */}
        <div className="w-[calc((100vw/3)*2)] h-16 mb-4 bg-zinc-900 rounded-md animate-pulse"></div>
        {/* VENUE */}
        <div className="flex gap-2 items-center animate-pulse">
          <div className="h-5 w-5 bg-zinc-900 rounded-full"></div>
          <div className="h-5 w-12 bg-zinc-900 rounded-md"></div>
          <div className="h-5 w-16 bg-zinc-900 rounded-md"></div>
        </div>
        {/* DATE */}
        <div className="flex gap-2 items-center animate-pulse">
          <div className="h-5 w-5 bg-zinc-900 rounded-full"></div>
          <div className="h-5 w-16 bg-zinc-900 rounded-md"></div>
          <div className="h-5 w-20 bg-zinc-900 rounded-md"></div>
        </div>
        {/* GENRES */}
        <div className="flex gap-2 items-center animate-pulse">
          <div className="h-5 w-5 bg-zinc-900 rounded-full"></div>
          <div className="h-5 w-12 bg-zinc-900 rounded-md"></div>
          <div className="h-5 w-10 bg-zinc-900 rounded-md"></div>
        </div>
      </div>
    </div>
  )
}
