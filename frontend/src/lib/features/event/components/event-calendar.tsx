import { differenceInMinutes, startOfHour, addHours, format } from "date-fns";

import { cn } from "@/lib/clsx";

import type { Event } from "@/lib/features/event/event";
import { useEffect, useMemo, useState, type HTMLAttributes } from "react";
import { Link } from "@tanstack/react-router";
import type { Concert } from "../concert";
import LinkButton from "@/lib/components/ui/button/link-button";
import { generateGoogleCalendarEventUrl } from "@/lib/google-calendar";
import { FaCalendarPlus } from "react-icons/fa6";

type Props = HTMLAttributes<HTMLDivElement> & {
  event: Event;
};

const EventCalendar = ({ event, ...rest }: Props) => {
  let concerts = event.concerts.sort(
    (a, b) => a.from.getTime() - b.from.getTime(),
  );

  const [calendarUrl, setCalendarUrl] = useState<URL | null>(null);

  useEffect(() => {
    const url = generateGoogleCalendarEventUrl(event);
    setCalendarUrl(url);
  }, [event.title, concerts.length]);

  let startHour =
    concerts.length > 0
      ? startOfHour(concerts[0].from)
      : startOfHour(new Date());

  let endHour =
    concerts.length > 0
      ? addHours(startOfHour(concerts[concerts.length - 1].to), 1)
      : addHours(startOfHour(new Date()), 4);

  let timeMarkers = useMemo(() => {
    let markers: Date[] = [];
    let current = startHour;

    while (current <= endHour) {
      markers = [...markers, current];
      current = addHours(current, 1);
    }

    return markers;
  }, [startHour, endHour]);

  let totalMinutes = Math.max(1, differenceInMinutes(endHour, startHour));

  return (
    <div className="flex w-full flex-col gap-8" {...rest}>
      <div className="flex items-center-safe justify-between">
        <div>
          <h3 className="mb-2 text-xl font-bold text-text-light">
            Program for {event.title}
          </h3>
          <span>{format(concerts[0].from, "EEEE, dd/MM/yyyy")}</span>
        </div>

        {calendarUrl && (
          <LinkButton
            className="hidden sm:flex"
            variant="secondary"
            to={calendarUrl.toString()}
            target="__blank"
          >
            <FaCalendarPlus />
            Føj til kalender
          </LinkButton>
        )}
      </div>

      <div className="overflow-y-scroll">
        <div className="grid h-full min-h-72 flex-1 gap-4 sm:grid-cols-[48px_1fr]">
          {/*  Timeline */}
          <div className="relative hidden sm:block">
            {timeMarkers.map((marker, i) => {
              const markerOffset = differenceInMinutes(marker, startHour);

              return (
                <div
                  key={marker.getTime()}
                  style={{ top: `calc(${markerOffset / totalMinutes} * 100%)` }}
                  className={cn("absolute flex w-full flex-col gap-1 text-sm", {
                    hidden: i === timeMarkers.length - 1,
                  })}
                >
                  <div className="h-[1px] w-full bg-zinc-800"></div>
                  <span>{format(marker, "HH:mm")}</span>
                </div>
              );
            })}
          </div>

          {/* Events */}
          <div className="relative">
            {event.concerts.map((concert) => (
              <Entry
                key={concert.id}
                concert={concert}
                totalMinutes={totalMinutes}
                startHour={startHour}
              />
            ))}
          </div>
        </div>

        {calendarUrl && (
          <LinkButton
            className="flex w-full sm:hidden"
            variant="secondary"
            to={calendarUrl.toString()}
            target="__blank"
          >
            <FaCalendarPlus />
            Føj til kalender
          </LinkButton>
        )}
      </div>
    </div>
  );
};

type EntryProps = {
  concert: Concert;
  totalMinutes: number;
  startHour: Date;
};

const Entry = ({ concert, totalMinutes, startHour }: EntryProps) => {
  const concertStartOffset = differenceInMinutes(concert.from, startHour);
  const concertDurationMinutes = differenceInMinutes(concert.to, concert.from);

  return (
    <Link
      style={{
        top: `calc(${concertStartOffset / totalMinutes} * 100%)`,
        height: `calc(${concertDurationMinutes / totalMinutes} * 100% - 1px)`,
      }}
      className="group absolute flex min-h-10 w-full justify-between overflow-hidden rounded-xl border border-t border-blue-800 bg-blue-950 p-4 text-sm transition-colors hover:border-blue-600 hover:bg-blue-900"
      to="/artists/$artistId"
      params={{ artistId: concert.artist.id.toString() }}
    >
      <p className="font-bold text-blue-200 group-hover:underline">
        {concert.artist.name}
      </p>
      <p className="text-blue-500">
        {format(concert.from, "HH:mm")} - {format(concert.to, "HH:mm")}
      </p>
    </Link>
  );
};

export default EventCalendar;
