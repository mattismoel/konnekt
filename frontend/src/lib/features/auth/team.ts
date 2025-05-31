import { idSchema, requestAndParse, type ID } from "@/lib/api"
import { createListResult, type ListResult } from "@/lib/query"
import { createUrl, type Query } from "@/lib/url"
import { z } from "zod"

export const teamTypes = z.union([
	z.literal("admin"),
	z.literal("member"),
	z.literal("event-management"),
	z.literal("booking"),
	z.literal("public-relations"),
	z.literal("visual-identity"),
	z.literal("team-management"),
	z.literal("project-leader"),
])

export type TeamType = z.infer<typeof teamTypes>

export const teamSchema = z.object({
	id: idSchema,
	name: teamTypes,
	displayName: z.string(),
	description: z.string()
})

export type Team = z.infer<typeof teamSchema>

export const memberTeams = async (memberId: ID): Promise<Team[]> => {
	const teams = await requestAndParse(
		createUrl(`/api/members/${memberId}/teams`),
		teamSchema.array(),
		"Could not fetch member teams",
	)

	return teams
}

export const listTeams = async (query?: Query): Promise<ListResult<Team>> => {
	const result = await requestAndParse(
		createUrl(`/api/teams`, query),
		createListResult(teamSchema),
		"Could not list teams",
	)

	return result
}

export const setMemberTeams = async (memberId: ID, teamIds: ID[]): Promise<void> => {
	await requestAndParse(
		createUrl(`/api/members/${memberId}/teams`),
		undefined,
		"Could not update member teams",
		{ body: teamIds, bodySchema: idSchema.array() },
		"PUT",
	)
}

/**
 * @description Checks whether or not a given member has all required input teams.
 */
export const hasAllTeams = (teams: Team[], reqTeamNames: TeamType[]): boolean => {
	return reqTeamNames.every(team => teams.some(t => t.name === team))
}
