import { useState } from "react";
import { type Artist } from "../artist";
import { useToast } from "@/lib/context/toast";
import { useQueryClient } from "@tanstack/react-query";
import { useAuth } from "@/lib/context/auth";
import List from "@/lib/components/list/list";
import { useSearch } from "@/lib/hooks/useSearch";
import Searchbar from "@/lib/components/searchbar";
import { Link } from "@tanstack/react-router";

type Props = {
  artists: Artist[];
  upcomingArtists: Artist[];
};

const ArtistList = ({ artists, upcomingArtists }: Props) => {
  const { search, setSearch, results } = useSearch(artists, "name");

  return (
    <>
      <Searchbar search={search} onChange={setSearch} />
      {search ? (
        <List>
          {results.map((artist) => (
            <Entry key={artist.id} artist={artist} />
          ))}
        </List>
      ) : (
        upcomingArtists.length > 0 && (
          <>
            <div>
              <p className="mb-4 font-medium">Kommende kunstnere</p>
              <List>
                {upcomingArtists.map((artist) => (
                  <Entry key={artist.id} artist={artist} />
                ))}
              </List>
            </div>
            <div>
              <p className="mb-4 font-medium">Alle kunstnere</p>
              <List>
                {artists.map((artist) => (
                  <Entry key={artist.id} artist={artist} />
                ))}
              </List>
            </div>
          </>
        )
      )}
    </>
  );
};

type EntryProps = {
  artist: Artist;
};

const Entry = ({ artist }: EntryProps) => {
  const { addToast } = useToast();
  const queryClient = useQueryClient();
  const { hasPermissions } = useAuth();
  let [showContextMenu, setShowContextMenu] = useState(false);

  return (
    <li className="group flex rounded-xl border border-zinc-800 bg-zinc-900 transition-colors hover:border-zinc-700 hover:bg-zinc-800">
      <Link
        to="/admin/artists/$artistId/edit"
        params={{ artistId: artist.id.toString() }}
        className="w-full py-4 pl-8"
      >
        <p className="font-medium group-hover:text-text-light">{artist.name}</p>
        <p className="text-base">
          {artist.genres.map((genre) => genre.name).join(", ")}
        </p>
      </Link>
    </li>
  );
};

export default ArtistList;
