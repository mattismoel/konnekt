import { PUBLIC_BACKEND_URL } from "$env/static/public"
import { requestAndParse } from "$lib/api"
import { createListResult, type ListResult } from "$lib/query"
import { createUrl, type Query } from "$lib/url"
import { z } from "zod"

export const teamTypes = z.union([
	z.literal("admin"),
	z.literal("member"),
	z.literal("event-management"),
	z.literal("booking"),
	z.literal("public-relations"),
	z.literal("visual-identity"),
	z.literal("team-management")
])

export type TeamType = z.infer<typeof teamTypes>

export const teamSchema = z.object({
	id: z.number().positive(),
	name: teamTypes,
	displayName: z.string(),
	description: z.string()
})

export type Team = z.infer<typeof teamSchema>

export const memberTeams = async (fetchFn: typeof fetch, memberId: number): Promise<Team[]> => {
	const teams = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/teams/${memberId}`),
		teamSchema.array(),
		"Could not fetch member teams",
	)

	return teams
}

export const listTeams = async (fetchFn: typeof fetch, query?: Query): Promise<ListResult<Team>> => {
	const result = await requestAndParse(
		fetchFn,
		createUrl(`${PUBLIC_BACKEND_URL}/teams`, query),
		createListResult(teamSchema),
		"Could not list teams",
	)

	return result
}


/**
 * @description Checks whether or not a given member has all required input teams.
 */
export const hasAllTeams = (teams: Team[], reqTeamNames: TeamType[]): boolean => {
	return reqTeamNames.every(team => teams.some(t => t.name === team))
}

export const hasSomeTeam = (teams: Team[], reqTeamNames: TeamType[]): boolean => {
	return reqTeamNames.some(team => teams.some(t => t.name === team))
}

