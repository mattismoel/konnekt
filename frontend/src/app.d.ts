// See https://svelte.dev/docs/kit/types#app.d.ts

import type { Permission, Role } from "$lib/features/auth/member";
import type { Member } from "$lib/features/auth/member";

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals { }
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export { };
