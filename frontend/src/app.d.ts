// See https://svelte.dev/docs/kit/types#app.d.ts

import type { Permission, Role } from "$lib/auth";
import type { User } from "$lib/auth";

// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: User | null
			roles: Role[]
			permissions: Permission[]
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export { };
