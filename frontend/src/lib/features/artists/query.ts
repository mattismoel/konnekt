import { queryOptions } from "@tanstack/react-query";
import { artist } from "./artist";
import { createListResult } from "@/lib/api/list";
import { throwAPIError } from "@/lib/api/error";

export const upcomingArtistsQueryOpts = () => queryOptions({
	queryKey: ["artists", "upcoming"],
	queryFn: async () => {
		const res = await fetch("/api/artists")
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		const result = createListResult(artist).parse(await res.json())
		return result
	}
})

export const previousArtistsQueryOpts = () => queryOptions({
	queryKey: ["artists", "previous"],
	queryFn: async () => {
		const res = await fetch("/api/artists")
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		const result = createListResult(artist).parse(await res.json())
		return result
	}
})
