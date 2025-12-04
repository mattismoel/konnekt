import type { Actions } from "./$types";

export const actions: Actions = {
	default: async ({ locals, request, params: { userId } }) => {
		const formData = await request.formData();
		await locals.pb.collection("users").update(userId, formData);
	}
};
