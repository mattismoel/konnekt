// See https://svelte.dev/docs/kit/types#app.d.ts

import type { Role } from "$lib/auth";
import type { User } from "$lib/user";

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: User | null
			roles: Role[]
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export { };
