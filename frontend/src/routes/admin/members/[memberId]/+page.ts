import { memberById } from "$lib/features/auth/member";
import { listTeams, memberTeams } from "$lib/features/auth/team";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ fetch, params }) => {
	const member = await memberById(fetch, parseInt(params.memberId))

	const { records: allTeams } = await listTeams(fetch)

	const teams = await memberTeams(fetch, member.id)

	return { member, teams, allTeams }
}
