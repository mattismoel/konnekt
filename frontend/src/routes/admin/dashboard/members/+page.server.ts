import { listMembers } from "$lib/features/auth/member";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch }) => {
	const { records: members } = await listMembers(fetch)
	const { records: pending } = await listMembers(fetch, {
		filter: ["active=false"]
	})

	return { members, pending }
}
