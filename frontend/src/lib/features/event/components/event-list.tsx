import { useState } from "react";

import EventEntry from "./event-list-entry";
import type { Event } from "../event";
import List from "@/lib/components/list/list";
import Button from "@/lib/components/ui/button/button";
import { useSearch } from "@/lib/hooks/useSearch";
import Searchbar from "@/lib/components/searchbar";

type Props = {
  previousEvents: Event[];
  upcomingEvents: Event[];
};

const EventList = ({ previousEvents, upcomingEvents }: Props) => {
  const [showPrevious, setShowPrevious] = useState(false);

  const { search, setSearch, results } = useSearch(
    [...previousEvents, ...upcomingEvents],
    "title",
  );

  return (
    <>
      <Searchbar search={search} onChange={setSearch} className="mb-8" />

      {search ? (
        <List>
          {results.map((event) => (
            <EventEntry key={event.id} event={event} />
          ))}
        </List>
      ) : (
        <>
          <div className="mb-8">
            <p className="mb-4 font-medium">Kommende events</p>
            <List>
              {upcomingEvents.map((event) => (
                <EventEntry key={event.id} event={event} />
              ))}
            </List>
          </div>

          {showPrevious && (
            <div className="mb-8">
              <p className="mb-4 font-medium">Tidligere events</p>
              <List>
                {previousEvents.map((event) => (
                  <EventEntry key={event.id} event={event} />
                ))}
              </List>
            </div>
          )}

          {previousEvents.length > 0 && (
            <Button
              onClick={() => setShowPrevious((prev) => !prev)}
              type="button"
              variant="secondary"
              className="w-full"
            >
              {showPrevious ? "Skjul tidligere events" : "Vis tidligere events"}
            </Button>
          )}
        </>
      )}
    </>
  );
};

export default EventList;
