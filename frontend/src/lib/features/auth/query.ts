import { queryOptions } from "@tanstack/react-query";
import { listMembers, memberById } from "./member";
import { listTeams, memberTeams } from "./team";
import type { ID } from "@/lib/api";

export const membersQueryOpts = queryOptions({
	queryKey: ["members"],
	queryFn: () => listMembers({
		filter: ["active=true"]
	})
})

export const pendingMembersQueryOpts = queryOptions({
	queryKey: ["members", "non-approved"],
	queryFn: () => listMembers({
		filter: ["active=false"]
	})
})

export const createMemberByIdQueryOpts = (memberId: ID) =>
	queryOptions({
		queryKey: ["members", memberId],
		queryFn: () => memberById(memberId)
	})

export const teamsQueryOpts = queryOptions({
	queryKey: ["teams"],
	queryFn: () => listTeams()
})

export const createMemberTeamsQueryOpts = (memberId: ID) => queryOptions({
	queryKey: ["members", "teams", memberId],
	queryFn: () => memberTeams(memberId)
})
