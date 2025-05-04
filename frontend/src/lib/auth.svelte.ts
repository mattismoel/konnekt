import type { Member } from "./features/auth/member"
import type { Permission } from "./features/auth/permission"
import type { Team } from "./features/auth/team"

type Auth = {
	member: Member | null
	permissions: Permission[]
	teams: Team[]
}

let auth = $state<Auth>({
	member: null,
	permissions: [],
	teams: []
})

export const authStore = {
	get permissions() {
		return auth?.permissions
	},
	get teams() {
		return auth.teams
	},
	get member() {
		return auth.member
	},
	set auth(newAuth: Auth) {
		auth = newAuth
	}
}
