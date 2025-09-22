import { createContext, useContext, useState, type PropsWithChildren } from "react"
import { memberSchema, uploadAvatar, type Member } from "../features/auth/member"
import type z from "zod";
import { loginFormSchema, memberPermissions, memberTeams, type Permission, type registerFormSchema, type Team } from "../features/auth/auth";
import { throwAPIError } from "../api/error";

export type AuthContext = {
	member: Member | undefined | null;
	teams: Team[] | undefined | null;
	perms: Permission[] | undefined | null;

	registerMember: (form: z.infer<typeof registerFormSchema>) => Promise<void>;
	loginMember: (form: z.infer<typeof loginFormSchema>) => Promise<void>;

	isAuthenticated: () => Promise<boolean>;
}

const AuthContext = createContext<AuthContext | null>(null)

export const useAuthContext = () => {
	const ctx = useContext(AuthContext)
	if (ctx === null) throw new Error("No AuthContext.Provider found")

	return ctx
}

export const AuthProvider = ({ children }: PropsWithChildren) => {
	const [member, setMember] = useState<Member | null | undefined>(undefined)
	const [perms, setPerms] = useState<Permission[] | undefined | null>(undefined)
	const [teams, setTeams] = useState<Team[] | undefined | null>(undefined)

	const invalidateAll = () => {
		setMember(null)
		setPerms(null)
		setTeams(null)
	}

	const registerMember = async (form: z.infer<typeof registerFormSchema>) => {
		try {
			const { avatarFile } = form
			const avatarUrl = await uploadAvatar(avatarFile)

			const res = await fetch("/api/auth/register", {
				body: JSON.stringify({ ...form, avatarUrl }),
				method: "POST"
			})

			if (!res.ok) {
				throwAPIError(await res.json())
			}

			const member = memberSchema.parse(await res.json())
			const { records: permissions } = await memberPermissions(member.id)
			const { records: teams } = await memberTeams(member.id)
			setMember(member)
			setPerms(permissions)
			setTeams(teams)
		} catch (e) {
			invalidateAll()
		}
	}

	const loginMember = async (form: z.infer<typeof loginFormSchema>) => {
		try {
			const res = await fetch("/api/auth/login", {
				method: "POST",
				body: JSON.stringify(form)
			})

			if (!res.ok) throwAPIError(await res.json())

			const member = memberSchema.parse(await res.json())
			const { records: permissions } = await memberPermissions(member.id)
			const { records: teams } = await memberTeams(member.id)
			setMember(member)
			setPerms(permissions)
			setTeams(teams)
		} catch (e) {
			invalidateAll()
		}
	}

	const isAuthenticated = async (): Promise<boolean> => {
		try {
			const res = await fetch("/api/session")
			if (!res.ok) {
				throwAPIError(await res.json())
			}
			return true
		} catch (e) {
			invalidateAll()
			return false
		}
	}

	return (
		<AuthContext.Provider value={{ member, perms, teams, registerMember, loginMember, isAuthenticated }}>
			{children}
		</AuthContext.Provider>
	)
}
