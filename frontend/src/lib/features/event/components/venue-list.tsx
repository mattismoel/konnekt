import { useMemo, useState } from "react"
import { deleteVenue, type Venue } from "../venue"
import SearchList from "@/lib/components/search-list"
import { useQueryClient } from "@tanstack/react-query"
import { useAuth } from "@/lib/context/auth"
import { useToast } from "@/lib/context/toast"
import { APIError } from "@/lib/api"
import List from "@/lib/components/list/list"
import ContextMenu from "@/lib/components/context-menu"

const filterVenuesBySearch = (venues: Venue[], query: string) => query
	? venues.filter(venue => venue.name.toLowerCase().includes(query.toLowerCase()))
	: venues

type Props = {
	venues: Venue[]
}

const VenueList = ({ venues }: Props) => {
	const [search, setSearch] = useState("")

	let filteredVenues = useMemo(() =>
		filterVenuesBySearch(venues, search),
		[search, venues]
	)

	return (
		<SearchList search={search} onChange={setSearch}>
			{filteredVenues.map(venue => <Entry key={venue.id} venue={venue} />)}
		</SearchList>
	)
}
type EntryProps = {
	venue: Venue;
};

const Entry = ({ venue }: EntryProps) => {
	const [showContextMenu, setShowContextMenu] = useState(false);
	const queryClient = useQueryClient()

	const { hasPermissions } = useAuth()
	const { addToast } = useToast()

	const handleDeleteVenue = async () => {
		if (!confirm(`Er sikker på, at du vil slette venue "${venue.name}"?`)) return;

		try {
			await deleteVenue(venue.id);
			addToast('Venue slettet');
			await queryClient.invalidateQueries({ queryKey: ["venues"] });
		} catch (e) {
			if (e instanceof APIError) {
				addToast('Kunne ikke slette venue', e.message, 'error');
				return;
			}

			addToast('Kunne ikke slette venue', 'Noget gik galt...', 'error');
			return;
		}
	};

	return (
		<List.Entry>
			<List.Entry.LinkSection to="/admin/venues/$venueId/edit" params={{ venueId: venue.id.toString() }}>
				<span>{venue.name}</span>
				<span className="text-text/50">{venue.city}, {venue.countryCode}</span>
			</List.Entry.LinkSection>

			<List.Entry.Section className="w-min">
				<ContextMenu.Button onClick={() => (setShowContextMenu(true))} />
			</List.Entry.Section>

			<ContextMenu show={showContextMenu} className="absolute top-1/2 right-4" onClose={() => setShowContextMenu(false)}>
				<ContextMenu.LinkEntry
					disabled={!hasPermissions(['edit:venue'])}
					to="/admin/venues/$venueId/edit"
					params={{ venueId: venue.id.toString() }}
				>
					Redigér
				</ContextMenu.LinkEntry>
				<ContextMenu.Entry
					disabled={!hasPermissions(['delete:venue'])}
					onClick={handleDeleteVenue}
				>
					Slet
				</ContextMenu.Entry>
			</ContextMenu>
		</List.Entry>
	)
}

export default VenueList
