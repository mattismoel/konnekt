import { memberById, memberSession } from "$lib/features/auth/member";
import { hasPermissions } from "$lib/features/auth/permission";
import { redirect } from "@sveltejs/kit";
import { listTeams } from "$lib/features/auth/team";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ fetch, params }) => {
	const currentMember = await memberSession(fetch)
	const member = await memberById(fetch, parseInt(params.memberId))

	const { records: teams } = await listTeams(fetch)

	if (!hasPermissions(currentMember.permissions, ["view:member"])) {
		throw redirect(302, "/admin/dashboard")
	}

	return {
		currentMember,
		member,
		teams
	}
}
