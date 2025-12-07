import { form, getRequestEvent, query } from "$app/server";
import { loginForm, registerForm, userSchema } from "./auth";
import { id } from "$lib/model";
import { redirect } from "@sveltejs/kit";

export const register = form(registerForm, async (data) => {
	const { locals } = getRequestEvent();

	const formData = new FormData();
	formData.append("email", data.email);
	formData.append("firstName", data.firstName);
	formData.append("lastName", data.lastName);
	formData.append("avatar", data.avatar);
	formData.append("password", data.password);
	formData.append("passwordConfirm", data.passwordConfirm);

	await locals.pb.collection("users").create(formData, {
		body: formData
	});
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
