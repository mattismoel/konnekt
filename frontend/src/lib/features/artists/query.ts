import { queryOptions } from "@tanstack/react-query";
import { artistSchema } from "./artist";
import { createListResult } from "@/lib/api/list";
import { throwAPIError } from "@/lib/api/error";

export const upcomingArtistsQueryOpts = () => queryOptions({
	queryKey: ["artists", "upcoming"],
	queryFn: async () => {
		const res = await fetch("/api/artists")
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		const result = createListResult(artistSchema).parse(await res.json())
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

		const result = createListResult(artistSchema).parse(await res.json())
		return result
	}
})

export const artistByIdQueryOpts = (artistId: number) => queryOptions({
	queryKey: ["artists", artistId],
	queryFn: async () => {
		const res = await fetch(`/api/artists/${artistId}`)
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		return artistSchema.parse(await res.json())
	}
})
