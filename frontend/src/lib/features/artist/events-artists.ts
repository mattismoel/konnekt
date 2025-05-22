import { removeDuplicates } from "../../array"
import type { Artist } from "./artist"
import type { Event } from "../event/event"

export const eventsArtists = (events: Event[]): Artist[] => {
	return removeDuplicates(
		events.flatMap(({ concerts }) => concerts).map(({ artist }) => artist)
	)
}
