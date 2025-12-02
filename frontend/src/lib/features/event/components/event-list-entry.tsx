import { format } from "date-fns";

import { DATETIME_FORMAT } from "@/lib/time";

import { type Event } from "../event";
import type { Artist } from "../../artist/artist";

import { Link } from "@tanstack/react-router";

type Props = {
  event: Event;
};

const formatArtists = (artists: Artist[]): string => {
  if (artists.length > 2) {
    return (
      artists
        .slice(0, 2)
        .map((a) => a.name)
        .join(", ") +
      " " +
      `(+${artists.length - 2} mere)`
    );
  }

  return artists.map((a) => a.name).join(", ");
};

const EventEntry = ({ event }: Props) => {
  let artists = event.concerts.map((concert) => concert.artist);

  return (
    <li className="group flex items-center rounded-xl border border-zinc-800 bg-zinc-900 transition-colors hover:border-zinc-700 hover:bg-zinc-800">
      <Link
        to="/admin/events/$eventId/edit"
        params={{ eventId: event.id.toString() }}
        className="w-full py-4 pl-8"
      >
        <div>
          <p className="font-medium transition-colors group-hover:text-text-light group-hover:underline">
            {event.title}
          </p>
          <div className="text-base">
            <p>{format(event.concerts[0].from, DATETIME_FORMAT)}</p>
            <p>{formatArtists(artists)}</p>
          </div>
        </div>
      </Link>
    </li>
  );
};

export default EventEntry;
