import { redirect } from "@sveltejs/kit";
import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = async ({ locals }) => {
	if (!locals.pb.authStore.isValid) {
		return redirect(302, "/auth/login");
	}
};
