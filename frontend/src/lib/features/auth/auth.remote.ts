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

export const getUser = query(id, async (memberId) => {
	const { locals } = getRequestEvent();
	const record = await locals.pb.collection("usersr").getOne(memberId);
	return userSchema.parse(record);
});

export const login = form(loginForm, async ({ email, password }) => {
	const { locals } = getRequestEvent();

	await locals.pb.collection("users").authWithPassword(email, password);
	return redirect(302, "/admin/dashboard");
});
