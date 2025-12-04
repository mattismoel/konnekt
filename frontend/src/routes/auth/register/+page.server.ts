import { registerForm, type RegisterForm } from "$lib/features/auth";
import { fail } from "@sveltejs/kit";
import type { Actions } from "./$types";
import { flattenError } from "zod";

export const actions: Actions = {
	default: async ({ locals, request }) => {
		const formData = await request.formData();

		const { success, error } = registerForm.safeParse(Object.fromEntries(formData));
		if (!success) {
			const { password, passwordConfirm, avatar, ...returnData } = {
				...(Object.fromEntries(formData) as RegisterForm)
			};

			return fail(400, {
				...flattenError(error),
				data: returnData
			});
		}

		await locals.pb.collection("users").create(formData);
	}
};
