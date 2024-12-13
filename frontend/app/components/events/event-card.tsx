import { useState } from "react";
import { EventDTO } from "@/lib/dto/event.dto";
import { formatDateString } from "@/lib/time";
import { cn } from "@/lib/utils";

type Props = {
  event: EventDTO | undefined;
  isLoading: boolean;
}

export const EventCard = ({ event, isLoading }: Props) => {
  const [mousePosX, setMousePosX] = useState(0);
  const [mousePosY, setMousePosY] = useState(0);

  const [isFocused, setIsFocused] = useState(false);

  const setMousePos = (x: number, y: number) => {
    setMousePosX(x);
    setMousePosY(y);
  }

  const toggleFocus = () => {
    setIsFocused(!isFocused)
  }

  if (isLoading) return <Skeleton />

  return (
    <a
      className={cn("group")}
      href={`/events/${event?.id}`}
      onMouseEnter={toggleFocus}
      onMouseLeave={toggleFocus}
    >
      <div
        role="none"
        className={"relative w-full h-64 overflow-hidden"}
        onMouseMove={(e) => {
          const rect = e.currentTarget.getBoundingClientRect();
          setMousePos(e.clientX - rect.left, e.clientY - rect.top);
        }}
      >
        <img
          src={event?.coverImageUrl}
          alt={event?.title}
          className="w-full h-full object-cover scale-110 md:brightness-90 group-hover:brightness-100 group-hover:scale-100 transition-all duration-200"
        />
        <div
          className="absolute bottom-0 left-0 h-1/2 w-full bg-gradient-to-t from-black/80 opacity-0 group-hover:opacity-100 transition-opacity duration-300"
        ></div>
        <div
          className="absolute bottom-0 left-0 h-full w-full border border-white/0 mix-blend-overlay transition-all group-hover:border-white/50"
        ></div>
        <div
          className="flex flex-col px-5 pb-5 absolute bottom-0 left-0 text-white md:translate-y-full md:group-hover:translate-y-0 transition-all duration-100"
        >
          <h3 className="font-bold text-3xl mb-2">{event?.title}</h3>
          <span>{formatDateString(event?.fromDate || new Date())}</span>
        </div>
        <div
          className={`absolute h-72 w-72 -translate-x-1/2  -translate-y-1/2 blur-3xl bg-white/50 mix-blend-overlay scale-0  pointer-events-none transition-transform duration-400 group-[.focused]:scale-100`}
          style={{ left: `${mousePosX}px`, top: `${mousePosY}px` }}
        ></div>
      </div>
    </a>
  )
}


const Skeleton = () => {
  return (
    <div
      className="relative w-full h-64 overflow-hidden bg-zinc-900 rounded-md animate-pulse"
    >
      <div
        className="w-full flex flex-col px-5 pb-5 absolute bottom-0 left-0"
      >
        <div
          className="h-8 mb-6 bg-zinc-800 rounded-md"
          style={{ width: `calc(100% * ${Math.random()})` }}
        ></div>
        <div className="h-4 w-24 bg-zinc-800 rounded-full"></div>
      </div>
    </div>
  )
}
