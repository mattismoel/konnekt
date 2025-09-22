import { queryOptions } from "@tanstack/react-query";
import { memberSchema } from "./member";
import { createListResult } from "@/lib/api/list";

export const membersQueryOpts = () => queryOptions({
	queryKey: ["teams", "all"],
	queryFn: async () => {
		const res = await fetch("/api/members")
		if (!res.ok) {
			throw new Error("Could not fetch members")
		}

		const members = createListResult(memberSchema).parse(await res.json())
		return members
	}
})
