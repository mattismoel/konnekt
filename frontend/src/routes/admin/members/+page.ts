import { listMembers } from "$lib/features/auth/member";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch }) => {
	const { records: members } = await listMembers(fetch)
	const { records: pending } = await listMembers(fetch, {
		filter: ["active=false"]
	})

	return { members, pending }
}
