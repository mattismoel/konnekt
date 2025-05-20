import { useMemo, useState } from "react";
import { deleteArtist, type Artist } from "../artist";
import SearchList from "@/lib/components/search-list";
import { useToast } from "@/lib/context/toast";
import { useQueryClient } from "@tanstack/react-query";
import { useAuth } from "@/lib/context/auth";
import { APIError } from "@/lib/api";
import List from "@/lib/components/list/list";
import ContextMenu from "@/lib/components/context-menu";

type Props = {
	artists: Artist[];
	upcomingArtists: Artist[];
};

const ArtistList = ({ artists, upcomingArtists }: Props) => {
	let [search, setSearch] = useState("");

	const filteredArtists = useMemo(() =>
		artists.filter((a) => a.name.toLowerCase().includes(search.toLowerCase())),
		[artists, search])

	return (
		<SearchList search={search} onChange={(newSearch) => setSearch(newSearch)}>
			{search
				? filteredArtists.map(artist => <Entry key={artist.id} artist={artist} />)
				: upcomingArtists.map(artist => <Entry key={artist.id} artist={artist} />)
			}

			<details>
				<summary className="mb-4">Alle kunstnere</summary>
				{artists.map(artist => <Entry key={artist.id} artist={artist} />)}
			</details>
		</SearchList >
	)
}


type EntryProps = {
	artist: Artist;
};

const Entry = ({ artist }: EntryProps) => {
	const { addToast } = useToast()
	const queryClient = useQueryClient()
	const { hasPermissions } = useAuth()
	let [showContextMenu, setShowContextMenu] = useState(false);

	const handleDelete = async () => {
		if (!confirm(`Er du sikke på, at du vil slette ${artist.name}?`)) return;

		try {
			await deleteArtist(artist.id);
			addToast('Kunstner slettet');
			queryClient.invalidateQueries({ queryKey: ["artists"] });
		} catch (e) {
			if (e instanceof APIError) {
				addToast('Kunne ikke slette kunstner', e.message, 'error');
				throw e;
			}

			addToast('Kunne ikke slette kunstner', 'Noget gik galt...', 'error');
			throw e;
		}
	};

	return (
		<List.Entry>
			<List.Entry.LinkSection to="/admin/artists/$artistId/edit" params={{ artistId: artist.id.toString() }}>
				<span>{artist.name}</span>
				<span className="text-text/50">{artist.genres.map((genre) => genre.name).join(', ')}</span>
			</List.Entry.LinkSection>

			<List.Entry.Section className="w-min" >
				<ContextMenu.Button onClick={() => setShowContextMenu(true)} />
			</List.Entry.Section>

			<ContextMenu show={showContextMenu} onClose={() => setShowContextMenu(false)} className="absolute top-1/2 right-4">
				<ContextMenu.LinkEntry
					disabled={!hasPermissions(['edit:artist'])}
					to="/admin/artists/$artistId/edit"
					params={{ artistId: artist.id.toString() }}
				>
					Redigér
				</ContextMenu.LinkEntry>
				<ContextMenu.Entry
					onClick={handleDelete}
					disabled={!hasPermissions(['delete:artist'])}
				>
					Slet
				</ContextMenu.Entry>
			</ContextMenu>
		</List.Entry>
	)
}


export default ArtistList
