import { type Venue } from "../venue";
import List from "@/lib/components/list/list";
import Searchbar from "@/lib/components/searchbar";
import { useSearch } from "@/lib/hooks/useSearch";
import { Link } from "@tanstack/react-router";

type Props = {
  venues: Venue[];
};

const VenueList = ({ venues }: Props) => {
  const { search, results, setSearch } = useSearch(venues, "name");

  return (
    <>
      <Searchbar search={search} onChange={setSearch} className="mb-4" />
      <List>
        {results.map((venue) => (
          <Entry key={venue.id} venue={venue} />
        ))}
      </List>
    </>
  );
};
type EntryProps = {
  venue: Venue;
};

const Entry = ({ venue }: EntryProps) => {
  return (
    <li className="group flex rounded-xl border border-zinc-800 bg-zinc-900 transition-colors hover:border-zinc-700 hover:bg-zinc-800">
      <Link
        to="/admin/venues/$venueId/edit"
        params={{ venueId: venue.id.toString() }}
        className="w-full py-4 pl-8"
      >
        <p className="font-medium transition-colors group-hover:text-text-light">
          {venue.name}
        </p>
        <p className="text-base">
          {venue.city}, {venue.countryCode}
        </p>
      </Link>
    </li>
  );
};

export default VenueList;
