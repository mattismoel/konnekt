import { form, getRequestEvent, query } from "$app/server";
import {
	loginForm,
	registerForm,
	teamSchema,
	memberSchema,
	teamType,
	permissionSchema,
	permissionType
} from "./auth";
import { id } from "$lib/model";
import { redirect } from "@sveltejs/kit";
import { queryOptions } from "$lib/query";
import { createFileUrl, type PBTeam, type PBUser } from "$lib/pocketbase";
import z from "zod";

export const register = form(registerForm, async (data) => {
	const { locals } = getRequestEvent();

	const { id: memberTeamId } = await locals.pb
		.collection("teams")
		.getFirstListItem<PBTeam>(locals.pb.filter("name = {:name}", { name: "member" }));

	await locals.pb.collection("users").create({ ...data, teams: [memberTeamId] });
});

export const getAuthenticatedMember = query(() => {
	const { locals } = getRequestEvent();

	if (!locals.pb.authStore.isValid || !locals.pb.authStore.record) return null;

	const member = getMember(locals.pb.authStore.record.id);
	return member;
});

export const hasPermissions = query(permissionType.array(), async (permissions) => {
	const member = await getAuthenticatedMember();
	if (!member) return false;

	const memberPermissions = member.teams.flatMap((teams) =>
		teams.permissions.flatMap((permission) => permission.name)
	);

	return permissions.every((permission) =>
		memberPermissions.some((memberPermission) => memberPermission === permission)
	);
});

export const getMember = query(id, async (memberId) => {
	const { locals } = getRequestEvent();
	const record = await locals.pb.collection("users").getOne<PBUser>(memberId, {
		expand: "teams,teams.permissions"
	});
	return memberSchema.parse({
		...record,
		avatar: record.avatar ? createFileUrl("users", record.id, record.avatar) : undefined,
		teams: record.expand?.teams.map((team) => ({ ...team, permissions: team.expand?.permissions }))
	});
});

export const login = form(loginForm, async ({ email, password }) => {
	const { locals } = getRequestEvent();

	await locals.pb.collection("users").authWithPassword(email, password);
	return redirect(302, "/admin/dashboard");
});

export const getMembers = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const { items, ...rest } = await locals.pb
		.collection("users")
		.getList<PBUser>(opts?.page, opts?.perPage, {
			sort: "firstName",
			expand: "teams,teams.permissions"
		});

	const members = memberSchema.array().parse(
		items.map((member) => ({
			...member,
			avatar: member.avatar ? createFileUrl("users", member.id, member.avatar) : undefined,
			teams: member.expand?.teams.map((team) => ({
				...team,
				permissions: team.expand?.permissions
			}))
		}))
	);

	return { items: members, ...rest };
});

export const getTeam = query(id, async (id) => {
	const { locals } = getRequestEvent();

	const team = locals.pb.collection("teams").getOne<PBTeam>(id);
	return teamSchema.parse(team);
});

export const getTeams = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const { items, ...rest } = await locals.pb
		.collection("teams")
		.getList<PBTeam>(opts?.page, opts?.perPage, {
			sort: "name",
			expand: "permissions"
		});

	const teams = teamSchema.array().parse(
		items.map((record) => ({
			...record,
			permissions: record.expand?.permissions
		}))
	);

	return { items: teams, ...rest };
});
