import type { z } from "zod"
import { login, logOut, register, type loginForm, type registerForm } from "../features/auth/auth"
import { memberSession, type Member } from "../features/auth/member";
import { memberPermissions, type Permission, type PermissionType } from "../features/auth/permission";
import { memberTeams, type Team, type TeamType } from "../features/auth/team";
import { createContext, useContext, useState, type PropsWithChildren } from "react";


type AuthContext = {
	handleLogin: (form: z.infer<typeof loginForm>) => void;
	handleRegister: (form: z.infer<typeof registerForm>) => void;
	handleLogout: () => void;

	refetch: () => void;

	hasPermissions: (perms: PermissionType[]) => boolean;
	isOnSomeTeam: (teams: TeamType[]) => boolean;

	member: Member | undefined | null
	permissions: Permission[]
	teams: Team[]
}

const AuthContext = createContext<AuthContext | undefined>(undefined)

export const AuthProvider = ({ children }: PropsWithChildren) => {
	const [member, setMember] = useState<Member | null>()
	const [permissions, setPermissions] = useState<Permission[]>([])
	const [teams, setTeams] = useState<Team[]>([])

	const handleLogin = async (form: z.infer<typeof loginForm>) => {
		try {
			const member = await login(form)
			const permissions = await memberPermissions(member.id)
			const teams = await memberTeams(member.id)

			setMember(member)
			setPermissions(permissions)
			setTeams(teams)
		} catch (e) {
			console.error(e)
			setMember(null)
			setPermissions([])
			setTeams([])
		}
	}

	const handleRegister = async (form: z.infer<typeof registerForm>) => {
		try {
			const member = await register(form)
			const permissions = await memberPermissions(member.id)
			const teams = await memberTeams(member.id)

			setMember(member)
			setPermissions(permissions)
			setTeams(teams)
		} catch (e) {
			setMember(null)
			setPermissions([])
			setTeams([])
		}
	}

	const handleLogout = async () => {
		try {
			await logOut()
			setMember(null)
			setPermissions([])
			setTeams([])
		} catch (e) {
			setMember(null)
			setPermissions([])
			setTeams([])
		}
	}

	const refetch = async () => {
		try {

			const member = await memberSession()
			const permissions = await memberPermissions(member.id)
			const teams = await memberTeams(member.id)

			setMember(member)
			setPermissions(permissions)
			setTeams(teams)
		} catch (e) {
			setMember(null)
			setPermissions([])
			setTeams([])
		}
	}

	const hasPermissions = (perms: PermissionType[]): boolean => {
		return perms.every(perm => permissions.some(p => p.name === perm))
	}

	const isOnSomeTeam = (ts: TeamType[]): boolean => {
		return ts.some(team => teams.some(t => t.name === team))
	}

	return (
		<AuthContext.Provider
			value={{
				handleLogin,
				handleRegister,
				handleLogout,
				refetch,
				hasPermissions,
				isOnSomeTeam,
				member,
				permissions,
				teams,
			}}
		>
			{children}
		</AuthContext.Provider>
	)
}

export const useAuth = () => {
	const authContext = useContext(AuthContext)
	if (!authContext) throw new Error("UserContext has no provider")

	return authContext
}
