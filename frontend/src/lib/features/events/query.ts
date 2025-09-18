import { throwAPIError } from "@/lib/api/error";
import { queryOptions } from "@tanstack/react-query";
import { event } from "./event";
import { createListResult } from "@/lib/api/list";

export const upcomingEventsQueryOpts = () => queryOptions({
	queryKey: ["events", "upcoming"],
	queryFn: async () => {
		const res = await fetch("/api/events")
		if (!res.ok) {
			throwAPIError(await res.json())
		}

		return createListResult(event).parse(await res.json())
	}
})
