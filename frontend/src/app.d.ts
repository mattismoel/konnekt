// See https://svelte.dev/docs/kit/types#app.d.ts

import type { Permission, Role } from "$lib/features/auth/user";
import type { User } from "$lib/features/auth/user";

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
