import type { Member } from "../features/auth/member"
import { type Team, type TeamType } from "../features/auth/team"

const includedTeamNames: TeamType[] = [
	"project-leader",
	"public-relations",
	"event-management",
	"booking",
	"visual-identity",
]

type Props = {
	allTeams: Team[]
	members: Member[]
}

const TeamDisplay = ({ members, allTeams }: Props) => {
	const includedTeams = allTeams.filter(team => includedTeamNames.includes(team.name))

	const includedMembers = members.filter(member => member.teams.some(team =>
		includedTeams.some(includedTeam => includedTeam.name === team.name)
	))

	return (
		<div className="@container flex flex-col gap-16">
			<div className="grid grid-cols-1 gap-8 @lg:grid-cols-2">
				{includedMembers.map(member => <MemberInfo key={member.id} member={member} includedTeams={includedTeams} />)}
			</div>
		</div>
	)
}

type MemberInfoProps = {
	member: Member
	includedTeams: Team[]
}

const MemberInfo = ({ member, includedTeams }: MemberInfoProps) => {
	const memberTeams = member.teams.filter(team => includedTeams.some(t => t.name === team.name))

	return (
		<div className="group bg-background flex flex-col border border-zinc-800 hover:border-zinc-700 rounded-sm overflow-hidden hover:bg-zinc-900 transition-colors">
			<div className="overflow-hidden">
				<img aria-disabled alt="Background blur" src={member.profilePictureUrl} loading="lazy" className="brightness-75 h-48 w-full object-cover scale-100 group-hover:scale-100 group-hover:brightness-100 transition-[scale,filter] duration-500" />
			</div>
			<div className="relative p-4 @lg:p-6 flex flex-col gap-2 cursor-default w-full object-cover">
				<img
					src={member.profilePictureUrl}
					alt={`${member.firstName} ${member.lastName}`}
					loading="lazy"
					className="pointer-events-none absolute h-full w-full blur-3xl opacity-0 group-hover:opacity-25 transition-opacity duration-1000"
				/>
				<span className="font-semibold">{member.firstName} {member.lastName}</span>

				<div className="flex flex-col text-text/50 text-sm">
					<span>{memberTeams.map(t => t.displayName).join(", ")}</span>
					<span>{member.email}</span>
				</div>
			</div>
		</div>
	)
}

export default TeamDisplay
